package services

import (
	"context"
	"fmt"
	fdsrunner "github.com/solo-io/gloo/projects/discovery/pkg/fds/runner"
	udsrunner "github.com/solo-io/gloo/projects/discovery/pkg/uds/runner"
	"github.com/solo-io/gloo/projects/gloo/pkg/runner"
	"net"
	"sync/atomic"
	"time"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"

	"github.com/solo-io/gloo/pkg/utils/settingsutil"

	"github.com/solo-io/gloo/pkg/utils/statusutils"

	"github.com/solo-io/gloo/projects/gateway/pkg/translator"

	"github.com/solo-io/gloo/projects/gloo/pkg/upstreams/consul"

	"github.com/solo-io/solo-kit/test/helpers"

	skkube "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"

	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"google.golang.org/grpc"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"k8s.io/client-go/kubernetes"
)

type TestClients struct {
	runner.ResourceClientset
	GatewayClient        gatewayv1.GatewayClient
	HttpGatewayClient    gatewayv1.MatchableHttpGatewayClient
	VirtualServiceClient gatewayv1.VirtualServiceClient
	ProxyClient          gloov1.ProxyClient
	UpstreamClient       gloov1.UpstreamClient
	SecretClient         gloov1.SecretClient
	ServiceClient        skkube.ServiceClient
	GlooPort             int
	RestXdsPort          int
}

// WriteSnapshot writes all resources in the ApiSnapshot to the cache
func (c TestClients) WriteSnapshot(ctx context.Context, snapshot *gloosnapshot.ApiSnapshot) error {
	// We intentionally create child resources first to avoid having the validation webhook reject
	// the parent resource

	writeOptions := clients.WriteOpts{
		Ctx:               ctx,
		OverwriteExisting: false,
	}
	for _, secret := range snapshot.Secrets {
		if _, writeErr := c.SecretClient.Write(secret, writeOptions); writeErr != nil {
			return writeErr
		}
	}
	for _, us := range snapshot.Upstreams {
		if _, writeErr := c.UpstreamClient.Write(us, writeOptions); writeErr != nil {
			return writeErr
		}
	}
	for _, vs := range snapshot.VirtualServices {
		if _, writeErr := c.VirtualServiceClient.Write(vs, writeOptions); writeErr != nil {
			return writeErr
		}
	}
	for _, hgw := range snapshot.HttpGateways {
		if _, writeErr := c.HttpGatewayClient.Write(hgw, writeOptions); writeErr != nil {
			return writeErr
		}
	}
	for _, gw := range snapshot.Gateways {
		if _, writeErr := c.GatewayClient.Write(gw, writeOptions); writeErr != nil {
			return writeErr
		}
	}
	for _, proxy := range snapshot.Proxies {
		if _, writeErr := c.ProxyClient.Write(proxy, writeOptions); writeErr != nil {
			return writeErr
		}
	}

	return nil
}

// DeleteSnapshot deletes all resources in the ApiSnapshot from the cache
func (c TestClients) DeleteSnapshot(ctx context.Context, snapshot *gloosnapshot.ApiSnapshot) error {
	// We intentionally delete resources in the reverse order that we create resources
	// If we delete child resources first, the validation webhook may reject the change

	deleteOptions := clients.DeleteOpts{
		Ctx:            ctx,
		IgnoreNotExist: true,
	}

	for _, gw := range snapshot.Gateways {
		gwNamespace, gwName := gw.GetMetadata().Ref().Strings()
		if deleteErr := c.GatewayClient.Delete(gwNamespace, gwName, deleteOptions); deleteErr != nil {
			return deleteErr
		}
	}
	for _, hgw := range snapshot.HttpGateways {
		hgwNamespace, hgwName := hgw.GetMetadata().Ref().Strings()
		if deleteErr := c.HttpGatewayClient.Delete(hgwNamespace, hgwName, deleteOptions); deleteErr != nil {
			return deleteErr
		}
	}
	for _, vs := range snapshot.VirtualServices {
		vsNamespace, vsName := vs.GetMetadata().Ref().Strings()
		if deleteErr := c.VirtualServiceClient.Delete(vsNamespace, vsName, deleteOptions); deleteErr != nil {
			return deleteErr
		}
	}
	for _, us := range snapshot.Upstreams {
		usNamespace, usName := us.GetMetadata().Ref().Strings()
		if deleteErr := c.UpstreamClient.Delete(usNamespace, usName, deleteOptions); deleteErr != nil {
			return deleteErr
		}
	}
	for _, secret := range snapshot.Secrets {
		secretNamespace, secretName := secret.GetMetadata().Ref().Strings()
		if deleteErr := c.SecretClient.Delete(secretNamespace, secretName, deleteOptions); deleteErr != nil {
			return deleteErr
		}
	}

	// Proxies are auto generated by Gateway resources
	// Therefore we delete Proxies after we have deleted the resources that may regenerate a Proxy
	for _, proxy := range snapshot.Proxies {
		proxyNamespace, proxyName := proxy.GetMetadata().Ref().Strings()
		if deleteErr := c.ProxyClient.Delete(proxyNamespace, proxyName, deleteOptions); deleteErr != nil {
			return deleteErr
		}
	}

	return nil
}

var glooPortBase = int32(30400)

func AllocateGlooPort() int32 {
	return atomic.AddInt32(&glooPortBase, 1) + int32(config.GinkgoConfig.ParallelNode*1000)
}

func RunGateway(ctx context.Context, justGloo bool) TestClients {
	ns := defaults.GlooSystem
	ro := &RunOptions{
		NsToWrite: ns,
		NsToWatch: []string{"default", ns},
		WhatToRun: What{
			DisableGateway: justGloo,
		},
		KubeClient: helpers.MustKubeClient(),
	}
	return RunGlooGatewayUdsFds(ctx, ro)
}

type What struct {
	DisableGateway bool
	DisableUds     bool
	DisableFds     bool
}

type RunOptions struct {
	NsToWrite        string
	NsToWatch        []string
	WhatToRun        What
	GlooPort         int32
	ValidationPort   int32
	RestXdsPort      int32
	Settings         *gloov1.Settings
	Cache            memory.InMemoryResourceCache
	KubeClient       kubernetes.Interface
	ConsulClient     consul.ConsulWatcher
	ConsulDnsAddress string
}

//noinspection GoUnhandledErrorResult
func RunGlooGatewayUdsFds(ctx context.Context, runOptions *RunOptions) TestClients {
	if runOptions.GlooPort == 0 {
		runOptions.GlooPort = AllocateGlooPort()
	}
	if runOptions.ValidationPort == 0 {
		runOptions.ValidationPort = AllocateGlooPort()
	}
	if runOptions.RestXdsPort == 0 {
		runOptions.RestXdsPort = AllocateGlooPort()
	}

	if runOptions.Cache == nil {
		runOptions.Cache = memory.NewInMemoryResourceCache()
	}

	settings := &gloov1.Settings{
		WatchNamespaces:    runOptions.NsToWatch,
		DiscoveryNamespace: runOptions.NsToWrite,
	}
	ctx = settingsutil.WithSettings(ctx, settings)

	if runOptions.Settings == nil {
		runOptions.Settings = settings
	}

	// Enable DevMode for tests
	runOptions.Settings.DevMode = true
	glooStartOpts := defaultGlooOpts(ctx, runOptions)

	glooStartOpts.ControlPlane.BindAddr.(*net.TCPAddr).Port = int(runOptions.GlooPort)
	glooStartOpts.ValidationServer.BindAddr.(*net.TCPAddr).Port = int(runOptions.ValidationPort)

	if glooStartOpts.Settings == nil {
		glooStartOpts.Settings = &gloov1.Settings{}
	}
	if glooStartOpts.Settings.GetGloo() == nil {
		glooStartOpts.Settings.Gloo = &gloov1.GlooOptions{}
	}
	if glooStartOpts.Settings.GetGloo().GetRestXdsBindAddr() == "" {
		glooStartOpts.Settings.GetGloo().RestXdsBindAddr = fmt.Sprintf("%s:%d", net.IPv4zero.String(), runOptions.RestXdsPort)
	}
	glooStartOpts.ControlPlane.StartGrpcServer = true
	glooStartOpts.ValidationServer.StartGrpcServer = true
	glooStartOpts.GatewayControllerEnabled = !runOptions.WhatToRun.DisableGateway

	go runner.StartGloo(glooStartOpts)

	if !runOptions.WhatToRun.DisableFds {
		go func() {
			defer GinkgoRecover()
			fdsrunner.StartFDS(glooStartOpts)
		}()
	}
	if !runOptions.WhatToRun.DisableUds {
		go func() {
			defer GinkgoRecover()
			udsrunner.StartUDS(glooStartOpts)
		}()
	}

	resourceClientset := glooStartOpts.ResourceClientset
	return TestClients{
		GatewayClient:        resourceClientset.Gateways,
		HttpGatewayClient:    resourceClientset.MatchableHttpGateways,
		VirtualServiceClient: resourceClientset.VirtualServices,
		UpstreamClient:       resourceClientset.Upstreams,
		SecretClient:         resourceClientset.Secrets,
		ProxyClient:          resourceClientset.Proxies,
		ServiceClient:        glooStartOpts.KubeServiceClient,
		GlooPort: int(runOptions.GlooPort),
		RestXdsPort: int(runOptions.RestXdsPort),
	}
}

func defaultGlooOpts(ctx context.Context, runOptions *RunOptions) runner.StartOpts {
	ctx = contextutils.WithLogger(ctx, "gloo")
	logger := contextutils.LoggerFrom(ctx)
	grpcServer := grpc.NewServer(grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zap.NewNop()),
			func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
				logger.Infof("gRPC call: %v", info.FullMethod)
				return handler(srv, ss)
			},
		)),
	)
	grpcServerValidation := grpc.NewServer(grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zap.NewNop()),
			func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
				logger.Infof("gRPC call: %v", info.FullMethod)
				return handler(srv, ss)
			},
		)),
	)

	var validationOpts *translator.ValidationOpts
	if runOptions.Settings.GetGateway().GetValidation().GetProxyValidationServerAddr() != "" {
		if validationOpts == nil {
			validationOpts = &translator.ValidationOpts{}
		}
		validationOpts.ProxyValidationServerAddress = runOptions.Settings.GetGateway().GetValidation().GetProxyValidationServerAddr()
	}
	if runOptions.Settings.GetGateway().GetValidation().GetAllowWarnings() != nil {
		if validationOpts == nil {
			validationOpts = &translator.ValidationOpts{}
		}
		validationOpts.AllowWarnings = runOptions.Settings.GetGateway().GetValidation().GetAllowWarnings().GetValue()
	}
	if runOptions.Settings.GetGateway().GetValidation().GetAlwaysAccept() != nil {
		if validationOpts == nil {
			validationOpts = &translator.ValidationOpts{}
		}
		validationOpts.AlwaysAcceptResources = runOptions.Settings.GetGateway().GetValidation().GetAlwaysAccept().GetValue()
	}

	// Generate a set of Gloo clients
	// Intentionally pass in a nil kubeCache to force the clients to be backed by memory
	runOptions.Settings.ConfigSource = nil
	glooClientset, typedClientset, err := runner.GenerateGlooClientsets(ctx, runOptions.Settings, nil, runOptions.Cache)
	if err != nil {
		panic("Something went wrong constructing in memory clients, this should never happen")
	}

	return runner.StartOpts{
		Settings:                runOptions.Settings,
		WriteNamespace:          runOptions.NsToWrite,
		StatusReporterNamespace: statusutils.GetStatusReporterNamespaceOrDefault(defaults.GlooSystem),
		ResourceClientset: glooClientset,

		KubeServiceClient:    typedClientset.KubeServiceClient,
		KubeClient:    typedClientset.KubeClient,
		KubeCoreCache: typedClientset.KubeCoreCache,

		WatchNamespaces:         runOptions.NsToWatch,
		WatchOpts: clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second / 10,
		},
		ControlPlane: runner.NewControlPlane(ctx, grpcServer, &net.TCPAddr{
			IP:   net.IPv4zero,
			Port: 8081,
		}, nil, true),
		ValidationServer: runner.NewValidationServer(ctx, grpcServerValidation, &net.TCPAddr{
			IP:   net.IPv4zero,
			Port: 8081,
		}, true),
		ProxyDebugServer: runner.NewProxyDebugServer(ctx, grpcServer, &net.TCPAddr{
			IP:   net.IPv4zero,
			Port: 8001,
		}, false),

		Consul: runner.ConsulStartOpts{
			ConsulWatcher: typedClientset.ConsulWatcher,
			DnsServer:     runOptions.ConsulDnsAddress,
		},
		GatewayControllerEnabled: true,
		ValidationOpts:           validationOpts,
	}
}

package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Contains static factory methods to create backends.
//
// Example:
//   var mesh mesh
//   var router virtualRouter
//   var service service
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
//   	Listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
//   			Port: jsii.Number(8080),
//   			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
//   				HealthyThreshold: jsii.Number(3),
//   				Interval: cdk.Duration_Seconds(jsii.Number(5)),
//   				Path: jsii.String("/ping"),
//   				Timeout: cdk.Duration_*Seconds(jsii.Number(2)),
//   				UnhealthyThreshold: jsii.Number(2),
//   			}),
//   			Timeout: &HttpTimeout{
//   				Idle: cdk.Duration_*Seconds(jsii.Number(5)),
//   			},
//   		}),
//   	},
//   	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
//   })
//
//   virtualService := appmesh.NewVirtualService(this, jsii.String("service-1"), &VirtualServiceProps{
//   	VirtualServiceProvider: appmesh.VirtualServiceProvider_VirtualRouter(router),
//   	VirtualServiceName: jsii.String("service1.domain.local"),
//   })
//
//   node.AddBackend(appmesh.Backend_VirtualService(virtualService))
//
// Experimental.
type Backend interface {
	// Return backend config.
	// Experimental.
	Bind(_scope awscdk.Construct) *BackendConfig
}

// The jsii proxy struct for Backend
type jsiiProxy_Backend struct {
	_ byte // padding
}

// Experimental.
func NewBackend_Override(b Backend) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.Backend",
		nil, // no parameters
		b,
	)
}

// Construct a Virtual Service backend.
// Experimental.
func Backend_VirtualService(virtualService IVirtualService, props *VirtualServiceBackendOptions) Backend {
	_init_.Initialize()

	if err := validateBackend_VirtualServiceParameters(virtualService, props); err != nil {
		panic(err)
	}
	var returns Backend

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.Backend",
		"virtualService",
		[]interface{}{virtualService, props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Backend) Bind(_scope awscdk.Construct) *BackendConfig {
	if err := b.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *BackendConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}


package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Contains static factory methods to create backends.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//   var router virtualRouter
//   var service service
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
//   	mesh: mesh,
//   	serviceDiscovery: appmesh.serviceDiscovery.cloudMap(service),
//   	listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
//   			port: jsii.Number(8080),
//   			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
//   				healthyThreshold: jsii.Number(3),
//   				interval: cdk.duration.seconds(jsii.Number(5)),
//   				path: jsii.String("/ping"),
//   				timeout: cdk.*duration.seconds(jsii.Number(2)),
//   				unhealthyThreshold: jsii.Number(2),
//   			}),
//   			timeout: &httpTimeout{
//   				idle: cdk.*duration.seconds(jsii.Number(5)),
//   			},
//   		}),
//   	},
//   	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
//   })
//
//   virtualService := appmesh.NewVirtualService(this, jsii.String("service-1"), &virtualServiceProps{
//   	virtualServiceProvider: appmesh.virtualServiceProvider.virtualRouter(router),
//   	virtualServiceName: jsii.String("service1.domain.local"),
//   })
//
//   node.addBackend(appmesh.backend.virtualService(virtualService))
//
type Backend interface {
	// Return backend config.
	Bind(_scope constructs.Construct) *BackendConfig
}

// The jsii proxy struct for Backend
type jsiiProxy_Backend struct {
	_ byte // padding
}

func NewBackend_Override(b Backend) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.Backend",
		nil, // no parameters
		b,
	)
}

// Construct a Virtual Service backend.
func Backend_VirtualService(virtualService IVirtualService, props *VirtualServiceBackendOptions) Backend {
	_init_.Initialize()

	if err := validateBackend_VirtualServiceParameters(virtualService, props); err != nil {
		panic(err)
	}
	var returns Backend

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Backend",
		"virtualService",
		[]interface{}{virtualService, props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Backend) Bind(_scope constructs.Construct) *BackendConfig {
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


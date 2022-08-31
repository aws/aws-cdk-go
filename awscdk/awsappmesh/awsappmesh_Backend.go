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
	var returns *BackendConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}


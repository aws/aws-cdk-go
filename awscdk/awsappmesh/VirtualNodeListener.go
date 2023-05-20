package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Defines listener for a VirtualNode.
//
// Example:
//   var mesh mesh
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
//   				Interval: awscdk.Duration_Seconds(jsii.Number(5)),
//   				Path: jsii.String("/ping"),
//   				Timeout: awscdk.Duration_*Seconds(jsii.Number(2)),
//   				UnhealthyThreshold: jsii.Number(2),
//   			}),
//   			Timeout: &HttpTimeout{
//   				Idle: awscdk.Duration_*Seconds(jsii.Number(5)),
//   			},
//   		}),
//   	},
//   	BackendDefaults: &BackendDefaults{
//   		TlsClientPolicy: &TlsClientPolicy{
//   			Validation: &TlsValidation{
//   				Trust: appmesh.TlsValidationTrust_File(jsii.String("/keys/local_cert_chain.pem")),
//   			},
//   		},
//   	},
//   	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
//   })
//
//   cdk.Tags_Of(node).Add(jsii.String("Environment"), jsii.String("Dev"))
//
type VirtualNodeListener interface {
	// Binds the current object when adding Listener to a VirtualNode.
	Bind(scope constructs.Construct) *VirtualNodeListenerConfig
}

// The jsii proxy struct for VirtualNodeListener
type jsiiProxy_VirtualNodeListener struct {
	_ byte // padding
}

func NewVirtualNodeListener_Override(v VirtualNodeListener) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualNodeListener",
		nil, // no parameters
		v,
	)
}

// Returns an GRPC Listener for a VirtualNode.
func VirtualNodeListener_Grpc(props *GrpcVirtualNodeListenerOptions) VirtualNodeListener {
	_init_.Initialize()

	if err := validateVirtualNodeListener_GrpcParameters(props); err != nil {
		panic(err)
	}
	var returns VirtualNodeListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualNodeListener",
		"grpc",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns an HTTP Listener for a VirtualNode.
func VirtualNodeListener_Http(props *HttpVirtualNodeListenerOptions) VirtualNodeListener {
	_init_.Initialize()

	if err := validateVirtualNodeListener_HttpParameters(props); err != nil {
		panic(err)
	}
	var returns VirtualNodeListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualNodeListener",
		"http",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns an HTTP2 Listener for a VirtualNode.
func VirtualNodeListener_Http2(props *Http2VirtualNodeListenerOptions) VirtualNodeListener {
	_init_.Initialize()

	if err := validateVirtualNodeListener_Http2Parameters(props); err != nil {
		panic(err)
	}
	var returns VirtualNodeListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualNodeListener",
		"http2",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns an TCP Listener for a VirtualNode.
func VirtualNodeListener_Tcp(props *TcpVirtualNodeListenerOptions) VirtualNodeListener {
	_init_.Initialize()

	if err := validateVirtualNodeListener_TcpParameters(props); err != nil {
		panic(err)
	}
	var returns VirtualNodeListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualNodeListener",
		"tcp",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNodeListener) Bind(scope constructs.Construct) *VirtualNodeListenerConfig {
	if err := v.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *VirtualNodeListenerConfig

	_jsii_.Invoke(
		v,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}


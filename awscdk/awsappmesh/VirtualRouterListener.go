package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the properties needed to define listeners for a VirtualRouter.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//
//   router := mesh.addVirtualRouter(jsii.String("router"), &VirtualRouterBaseProps{
//   	Listeners: []virtualRouterListener{
//   		appmesh.*virtualRouterListener_Http(jsii.Number(8080)),
//   	},
//   })
//
type VirtualRouterListener interface {
	// Called when the VirtualRouterListener type is initialized.
	//
	// Can be used to enforce
	// mutual exclusivity.
	Bind(scope constructs.Construct) *VirtualRouterListenerConfig
}

// The jsii proxy struct for VirtualRouterListener
type jsiiProxy_VirtualRouterListener struct {
	_ byte // padding
}

func NewVirtualRouterListener_Override(v VirtualRouterListener) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualRouterListener",
		nil, // no parameters
		v,
	)
}

// Returns a GRPC Listener for a VirtualRouter.
func VirtualRouterListener_Grpc(port *float64) VirtualRouterListener {
	_init_.Initialize()

	var returns VirtualRouterListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouterListener",
		"grpc",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// Returns an HTTP Listener for a VirtualRouter.
func VirtualRouterListener_Http(port *float64) VirtualRouterListener {
	_init_.Initialize()

	var returns VirtualRouterListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouterListener",
		"http",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// Returns an HTTP2 Listener for a VirtualRouter.
func VirtualRouterListener_Http2(port *float64) VirtualRouterListener {
	_init_.Initialize()

	var returns VirtualRouterListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouterListener",
		"http2",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// Returns a TCP Listener for a VirtualRouter.
func VirtualRouterListener_Tcp(port *float64) VirtualRouterListener {
	_init_.Initialize()

	var returns VirtualRouterListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouterListener",
		"tcp",
		[]interface{}{port},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualRouterListener) Bind(scope constructs.Construct) *VirtualRouterListenerConfig {
	if err := v.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *VirtualRouterListenerConfig

	_jsii_.Invoke(
		v,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}


package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v3"
)

// Represents the properties needed to define the provider for a VirtualService.
//
// Example:
//   var mesh mesh
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
//   	mesh: mesh,
//   	serviceDiscovery: appmesh.serviceDiscovery.dns(jsii.String("node")),
//   })
//
//   virtualService := appmesh.NewVirtualService(this, jsii.String("service-1"), &virtualServiceProps{
//   	virtualServiceProvider: appmesh.virtualServiceProvider.virtualNode(node),
//   	virtualServiceName: jsii.String("service1.domain.local"),
//   })
//
//   node.addBackend(appmesh.backend.virtualService(virtualService))
//
// Experimental.
type VirtualServiceProvider interface {
	// Enforces mutual exclusivity for VirtualService provider types.
	// Experimental.
	Bind(_construct constructs.Construct) *VirtualServiceProviderConfig
}

// The jsii proxy struct for VirtualServiceProvider
type jsiiProxy_VirtualServiceProvider struct {
	_ byte // padding
}

// Experimental.
func NewVirtualServiceProvider_Override(v VirtualServiceProvider) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.VirtualServiceProvider",
		nil, // no parameters
		v,
	)
}

// Returns an Empty Provider for a VirtualService.
//
// This provides no routing capabilities
// and should only be used as a placeholder.
// Experimental.
func VirtualServiceProvider_None(mesh IMesh) VirtualServiceProvider {
	_init_.Initialize()

	if err := validateVirtualServiceProvider_NoneParameters(mesh); err != nil {
		panic(err)
	}
	var returns VirtualServiceProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualServiceProvider",
		"none",
		[]interface{}{mesh},
		&returns,
	)

	return returns
}

// Returns a VirtualNode based Provider for a VirtualService.
// Experimental.
func VirtualServiceProvider_VirtualNode(virtualNode IVirtualNode) VirtualServiceProvider {
	_init_.Initialize()

	if err := validateVirtualServiceProvider_VirtualNodeParameters(virtualNode); err != nil {
		panic(err)
	}
	var returns VirtualServiceProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualServiceProvider",
		"virtualNode",
		[]interface{}{virtualNode},
		&returns,
	)

	return returns
}

// Returns a VirtualRouter based Provider for a VirtualService.
// Experimental.
func VirtualServiceProvider_VirtualRouter(virtualRouter IVirtualRouter) VirtualServiceProvider {
	_init_.Initialize()

	if err := validateVirtualServiceProvider_VirtualRouterParameters(virtualRouter); err != nil {
		panic(err)
	}
	var returns VirtualServiceProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.VirtualServiceProvider",
		"virtualRouter",
		[]interface{}{virtualRouter},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualServiceProvider) Bind(_construct constructs.Construct) *VirtualServiceProviderConfig {
	if err := v.validateBindParameters(_construct); err != nil {
		panic(err)
	}
	var returns *VirtualServiceProviderConfig

	_jsii_.Invoke(
		v,
		"bind",
		[]interface{}{_construct},
		&returns,
	)

	return returns
}


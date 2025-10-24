package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the properties needed to define the provider for a VirtualService.
//
// Example:
//   var mesh Mesh
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_Dns(jsii.String("node")),
//   })
//
//   virtualService := appmesh.NewVirtualService(this, jsii.String("service-1"), &VirtualServiceProps{
//   	VirtualServiceProvider: appmesh.VirtualServiceProvider_VirtualNode(node),
//   	VirtualServiceName: jsii.String("service1.domain.local"),
//   })
//
//   node.AddBackend(appmesh.Backend_VirtualService(virtualService))
//
type VirtualServiceProvider interface {
	// Enforces mutual exclusivity for VirtualService provider types.
	Bind(_construct constructs.Construct) *VirtualServiceProviderConfig
}

// The jsii proxy struct for VirtualServiceProvider
type jsiiProxy_VirtualServiceProvider struct {
	_ byte // padding
}

func NewVirtualServiceProvider_Override(v VirtualServiceProvider) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualServiceProvider",
		nil, // no parameters
		v,
	)
}

// Returns an Empty Provider for a VirtualService.
//
// This provides no routing capabilities
// and should only be used as a placeholder.
func VirtualServiceProvider_None(mesh IMesh) VirtualServiceProvider {
	_init_.Initialize()

	if err := validateVirtualServiceProvider_NoneParameters(mesh); err != nil {
		panic(err)
	}
	var returns VirtualServiceProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualServiceProvider",
		"none",
		[]interface{}{mesh},
		&returns,
	)

	return returns
}

// Returns a VirtualNode based Provider for a VirtualService.
func VirtualServiceProvider_VirtualNode(virtualNode IVirtualNode) VirtualServiceProvider {
	_init_.Initialize()

	if err := validateVirtualServiceProvider_VirtualNodeParameters(virtualNode); err != nil {
		panic(err)
	}
	var returns VirtualServiceProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualServiceProvider",
		"virtualNode",
		[]interface{}{virtualNode},
		&returns,
	)

	return returns
}

// Returns a VirtualRouter based Provider for a VirtualService.
func VirtualServiceProvider_VirtualRouter(virtualRouter IVirtualRouter) VirtualServiceProvider {
	_init_.Initialize()

	if err := validateVirtualServiceProvider_VirtualRouterParameters(virtualRouter); err != nil {
		panic(err)
	}
	var returns VirtualServiceProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualServiceProvider",
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


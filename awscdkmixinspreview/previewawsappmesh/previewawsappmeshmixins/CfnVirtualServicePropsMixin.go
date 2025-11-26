package previewawsappmeshmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappmesh/previewawsappmeshmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a virtual service within a service mesh.
//
// A virtual service is an abstraction of a real service that is provided by a virtual node directly or indirectly by means of a virtual router. Dependent services call your virtual service by its `virtualServiceName` , and those requests are routed to the virtual node or virtual router that is specified as the provider for the virtual service.
//
// For more information about virtual services, see [Virtual services](https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_services.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVirtualServicePropsMixin := awscdkmixinspreview.Mixins.NewCfnVirtualServicePropsMixin(&CfnVirtualServiceMixinProps{
//   	MeshName: jsii.String("meshName"),
//   	MeshOwner: jsii.String("meshOwner"),
//   	Spec: &VirtualServiceSpecProperty{
//   		Provider: &VirtualServiceProviderProperty{
//   			VirtualNode: &VirtualNodeServiceProviderProperty{
//   				VirtualNodeName: jsii.String("virtualNodeName"),
//   			},
//   			VirtualRouter: &VirtualRouterServiceProviderProperty{
//   				VirtualRouterName: jsii.String("virtualRouterName"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VirtualServiceName: jsii.String("virtualServiceName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html
//
type CfnVirtualServicePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnVirtualServiceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVirtualServicePropsMixin
type jsiiProxy_CfnVirtualServicePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnVirtualServicePropsMixin) Props() *CfnVirtualServiceMixinProps {
	var returns *CfnVirtualServiceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualServicePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppMesh::VirtualService`.
func NewCfnVirtualServicePropsMixin(props *CfnVirtualServiceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnVirtualServicePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVirtualServicePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVirtualServicePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnVirtualServicePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppMesh::VirtualService`.
func NewCfnVirtualServicePropsMixin_Override(c CfnVirtualServicePropsMixin, props *CfnVirtualServiceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnVirtualServicePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnVirtualServicePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVirtualServicePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnVirtualServicePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVirtualServicePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appmesh.mixins.CfnVirtualServicePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVirtualServicePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVirtualServicePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}


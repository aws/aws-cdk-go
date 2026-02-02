package previewawsappstreammixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappstream/previewawsappstreammixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AppStream::ImageBuilder` resource creates an image builder for Amazon AppStream 2.0. An image builder is a virtual machine that is used to create an image.
//
// The initial state of the image builder is `PENDING` . When it is ready, the state is `RUNNING` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnImageBuilderPropsMixin := awscdkmixinspreview.Mixins.NewCfnImageBuilderPropsMixin(&CfnImageBuilderMixinProps{
//   	AccessEndpoints: []interface{}{
//   		&AccessEndpointProperty{
//   			EndpointType: jsii.String("endpointType"),
//   			VpceId: jsii.String("vpceId"),
//   		},
//   	},
//   	AppstreamAgentVersion: jsii.String("appstreamAgentVersion"),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	DomainJoinInfo: &DomainJoinInfoProperty{
//   		DirectoryName: jsii.String("directoryName"),
//   		OrganizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   	},
//   	EnableDefaultInternetAccess: jsii.Boolean(false),
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	ImageArn: jsii.String("imageArn"),
//   	ImageName: jsii.String("imageName"),
//   	InstanceType: jsii.String("instanceType"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html
//
type CfnImageBuilderPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnImageBuilderMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnImageBuilderPropsMixin
type jsiiProxy_CfnImageBuilderPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnImageBuilderPropsMixin) Props() *CfnImageBuilderMixinProps {
	var returns *CfnImageBuilderMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageBuilderPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppStream::ImageBuilder`.
func NewCfnImageBuilderPropsMixin(props *CfnImageBuilderMixinProps, options *mixins.CfnPropertyMixinOptions) CfnImageBuilderPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnImageBuilderPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnImageBuilderPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appstream.mixins.CfnImageBuilderPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppStream::ImageBuilder`.
func NewCfnImageBuilderPropsMixin_Override(c CfnImageBuilderPropsMixin, props *CfnImageBuilderMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appstream.mixins.CfnImageBuilderPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnImageBuilderPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnImageBuilderPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appstream.mixins.CfnImageBuilderPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnImageBuilderPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appstream.mixins.CfnImageBuilderPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnImageBuilderPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnImageBuilderPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


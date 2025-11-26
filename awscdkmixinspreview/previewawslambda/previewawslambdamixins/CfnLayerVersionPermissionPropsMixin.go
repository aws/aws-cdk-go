package previewawslambdamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslambda/previewawslambdamixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Lambda::LayerVersionPermission` resource adds permissions to the resource-based policy of a version of an [Lambda layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) . Use this action to grant layer usage permission to other accounts. You can grant permission to a single account, all AWS accounts, or all accounts in an organization.
//
// > Since the release of the [UpdateReplacePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html) both `UpdateReplacePolicy` and `DeletionPolicy` are required to protect your Resources/LayerPermissions from deletion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLayerVersionPermissionPropsMixin := awscdkmixinspreview.Mixins.NewCfnLayerVersionPermissionPropsMixin(&CfnLayerVersionPermissionMixinProps{
//   	Action: jsii.String("action"),
//   	LayerVersionArn: jsii.String("layerVersionArn"),
//   	OrganizationId: jsii.String("organizationId"),
//   	Principal: jsii.String("principal"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html
//
type CfnLayerVersionPermissionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLayerVersionPermissionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLayerVersionPermissionPropsMixin
type jsiiProxy_CfnLayerVersionPermissionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLayerVersionPermissionPropsMixin) Props() *CfnLayerVersionPermissionMixinProps {
	var returns *CfnLayerVersionPermissionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersionPermissionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lambda::LayerVersionPermission`.
func NewCfnLayerVersionPermissionPropsMixin(props *CfnLayerVersionPermissionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLayerVersionPermissionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLayerVersionPermissionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLayerVersionPermissionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnLayerVersionPermissionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lambda::LayerVersionPermission`.
func NewCfnLayerVersionPermissionPropsMixin_Override(c CfnLayerVersionPermissionPropsMixin, props *CfnLayerVersionPermissionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnLayerVersionPermissionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLayerVersionPermissionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLayerVersionPermissionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnLayerVersionPermissionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLayerVersionPermissionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnLayerVersionPermissionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLayerVersionPermissionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLayerVersionPermissionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


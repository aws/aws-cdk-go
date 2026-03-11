package awscloudformation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CloudFormation::ResourceDefaultVersion` resource specifies the default version of a resource.
//
// The default version of a resource will be used in CloudFormation operations.
//
// For information about the CloudFormation registry, see [Managing extensions with the CloudFormation registry](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html) in the *CloudFormation User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnResourceDefaultVersionPropsMixin := awscdkcfnpropertymixins.Aws_cloudformation.NewCfnResourceDefaultVersionPropsMixin(&CfnResourceDefaultVersionMixinProps{
//   	TypeName: jsii.String("typeName"),
//   	TypeVersionArn: jsii.String("typeVersionArn"),
//   	VersionId: jsii.String("versionId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourcedefaultversion.html
//
type CfnResourceDefaultVersionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnResourceDefaultVersionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResourceDefaultVersionPropsMixin
type jsiiProxy_CfnResourceDefaultVersionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnResourceDefaultVersionPropsMixin) Props() *CfnResourceDefaultVersionMixinProps {
	var returns *CfnResourceDefaultVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceDefaultVersionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFormation::ResourceDefaultVersion`.
func NewCfnResourceDefaultVersionPropsMixin(props *CfnResourceDefaultVersionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnResourceDefaultVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResourceDefaultVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResourceDefaultVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudformation.CfnResourceDefaultVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFormation::ResourceDefaultVersion`.
func NewCfnResourceDefaultVersionPropsMixin_Override(c CfnResourceDefaultVersionPropsMixin, props *CfnResourceDefaultVersionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudformation.CfnResourceDefaultVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnResourceDefaultVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResourceDefaultVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cloudformation.CfnResourceDefaultVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourceDefaultVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cloudformation.CfnResourceDefaultVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourceDefaultVersionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnResourceDefaultVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


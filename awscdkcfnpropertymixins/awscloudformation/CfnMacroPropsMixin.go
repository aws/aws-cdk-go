package awscloudformation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CloudFormation::Macro` resource is a CloudFormation resource type that creates a CloudFormation macro to perform custom processing on CloudFormation templates.
//
// For more information, see [Perform custom processing on CloudFormation templates with template macros](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html) in the *CloudFormation User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMacroPropsMixin := awscdkcfnpropertymixins.Aws_cloudformation.NewCfnMacroPropsMixin(&CfnMacroMixinProps{
//   	Description: jsii.String("description"),
//   	FunctionName: jsii.String("functionName"),
//   	LogGroupName: jsii.String("logGroupName"),
//   	LogRoleArn: jsii.String("logRoleArn"),
//   	Name: jsii.String("name"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-macro.html
//
type CfnMacroPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMacroMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMacroPropsMixin
type jsiiProxy_CfnMacroPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMacroPropsMixin) Props() *CfnMacroMixinProps {
	var returns *CfnMacroMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMacroPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFormation::Macro`.
func NewCfnMacroPropsMixin(props *CfnMacroMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMacroPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMacroPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMacroPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudformation.CfnMacroPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFormation::Macro`.
func NewCfnMacroPropsMixin_Override(c CfnMacroPropsMixin, props *CfnMacroMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudformation.CfnMacroPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMacroPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMacroPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cloudformation.CfnMacroPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMacroPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cloudformation.CfnMacroPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMacroPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMacroPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


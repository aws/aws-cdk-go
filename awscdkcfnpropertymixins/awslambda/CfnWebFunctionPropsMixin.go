package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Lambda::WebFunction.
//
// Creates a Lambda Web function container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnWebFunctionPropsMixin := awscdkcfnpropertymixins.Aws_lambda.NewCfnWebFunctionPropsMixin(&CfnWebFunctionMixinProps{
//   	FunctionName: jsii.String("functionName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunction.html
//
type CfnWebFunctionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnWebFunctionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWebFunctionPropsMixin
type jsiiProxy_CfnWebFunctionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnWebFunctionPropsMixin) Props() *CfnWebFunctionMixinProps {
	var returns *CfnWebFunctionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebFunctionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lambda::WebFunction`.
func NewCfnWebFunctionPropsMixin(props *CfnWebFunctionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnWebFunctionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWebFunctionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWebFunctionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnWebFunctionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lambda::WebFunction`.
func NewCfnWebFunctionPropsMixin_Override(c CfnWebFunctionPropsMixin, props *CfnWebFunctionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnWebFunctionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnWebFunctionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWebFunctionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnWebFunctionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWebFunctionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnWebFunctionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWebFunctionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnWebFunctionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


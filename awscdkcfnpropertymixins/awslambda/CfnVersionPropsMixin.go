package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Lambda::Version` resource creates a [version](https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html) from the current code and configuration of a function. Use versions to create a snapshot of your function code and configuration that doesn't change.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnVersionPropsMixin := awscdkcfnpropertymixins.Aws_lambda.NewCfnVersionPropsMixin(&CfnVersionMixinProps{
//   	CodeSha256: jsii.String("codeSha256"),
//   	Description: jsii.String("description"),
//   	FunctionName: jsii.String("functionName"),
//   	FunctionScalingConfig: &FunctionScalingConfigProperty{
//   		MaxExecutionEnvironments: jsii.Number(123),
//   		MinExecutionEnvironments: jsii.Number(123),
//   	},
//   	ProvisionedConcurrencyConfig: &ProvisionedConcurrencyConfigurationProperty{
//   		ProvisionedConcurrentExecutions: jsii.Number(123),
//   	},
//   	RuntimePolicy: &RuntimePolicyProperty{
//   		RuntimeVersionArn: jsii.String("runtimeVersionArn"),
//   		UpdateRuntimeOn: jsii.String("updateRuntimeOn"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html
//
type CfnVersionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnVersionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVersionPropsMixin
type jsiiProxy_CfnVersionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnVersionPropsMixin) Props() *CfnVersionMixinProps {
	var returns *CfnVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVersionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lambda::Version`.
func NewCfnVersionPropsMixin(props *CfnVersionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lambda::Version`.
func NewCfnVersionPropsMixin_Override(c CfnVersionPropsMixin, props *CfnVersionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVersionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


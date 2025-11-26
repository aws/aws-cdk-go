package previewawslambdamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslambda/previewawslambdamixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Lambda::EventInvokeConfig` resource configures options for [asynchronous invocation](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html) on a version or an alias.
//
// By default, Lambda retries an asynchronous invocation twice if the function returns an error. It retains events in a queue for up to six hours. When an event fails all processing attempts or stays in the asynchronous invocation queue for too long, Lambda discards it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventInvokeConfigPropsMixin := awscdkmixinspreview.Mixins.NewCfnEventInvokeConfigPropsMixin(&CfnEventInvokeConfigMixinProps{
//   	DestinationConfig: &DestinationConfigProperty{
//   		OnFailure: &OnFailureProperty{
//   			Destination: jsii.String("destination"),
//   		},
//   		OnSuccess: &OnSuccessProperty{
//   			Destination: jsii.String("destination"),
//   		},
//   	},
//   	FunctionName: jsii.String("functionName"),
//   	MaximumEventAgeInSeconds: jsii.Number(123),
//   	MaximumRetryAttempts: jsii.Number(123),
//   	Qualifier: jsii.String("qualifier"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventinvokeconfig.html
//
type CfnEventInvokeConfigPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEventInvokeConfigMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEventInvokeConfigPropsMixin
type jsiiProxy_CfnEventInvokeConfigPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEventInvokeConfigPropsMixin) Props() *CfnEventInvokeConfigMixinProps {
	var returns *CfnEventInvokeConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventInvokeConfigPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lambda::EventInvokeConfig`.
func NewCfnEventInvokeConfigPropsMixin(props *CfnEventInvokeConfigMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEventInvokeConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEventInvokeConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEventInvokeConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventInvokeConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lambda::EventInvokeConfig`.
func NewCfnEventInvokeConfigPropsMixin_Override(c CfnEventInvokeConfigPropsMixin, props *CfnEventInvokeConfigMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventInvokeConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEventInvokeConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventInvokeConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventInvokeConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventInvokeConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventInvokeConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEventInvokeConfigPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnEventInvokeConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


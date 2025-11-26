package previewawslambdamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslambda/previewawslambdamixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Lambda::Alias` resource creates an [alias](https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html) for a Lambda function version. Use aliases to provide clients with a function identifier that you can update to invoke a different version.
//
// You can also map an alias to split invocation requests between two versions. Use the `RoutingConfig` parameter to specify a second version and the percentage of invocation requests that it receives.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAliasPropsMixin := awscdkmixinspreview.Mixins.NewCfnAliasPropsMixin(&CfnAliasMixinProps{
//   	Description: jsii.String("description"),
//   	FunctionName: jsii.String("functionName"),
//   	FunctionVersion: jsii.String("functionVersion"),
//   	Name: jsii.String("name"),
//   	ProvisionedConcurrencyConfig: &ProvisionedConcurrencyConfigurationProperty{
//   		ProvisionedConcurrentExecutions: jsii.Number(123),
//   	},
//   	RoutingConfig: &AliasRoutingConfigurationProperty{
//   		AdditionalVersionWeights: []interface{}{
//   			&VersionWeightProperty{
//   				FunctionVersion: jsii.String("functionVersion"),
//   				FunctionWeight: jsii.Number(123),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html
//
type CfnAliasPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAliasMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAliasPropsMixin
type jsiiProxy_CfnAliasPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAliasPropsMixin) Props() *CfnAliasMixinProps {
	var returns *CfnAliasMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAliasPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lambda::Alias`.
func NewCfnAliasPropsMixin(props *CfnAliasMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAliasPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAliasPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAliasPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnAliasPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lambda::Alias`.
func NewCfnAliasPropsMixin_Override(c CfnAliasPropsMixin, props *CfnAliasMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnAliasPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAliasPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAliasPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnAliasPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAliasPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnAliasPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAliasPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAliasPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


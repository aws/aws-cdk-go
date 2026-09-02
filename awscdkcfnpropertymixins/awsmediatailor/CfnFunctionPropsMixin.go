package awsmediatailor

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::MediaTailor::Function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnFunctionPropsMixin := awscdkcfnpropertymixins.Aws_mediatailor.NewCfnFunctionPropsMixin(&CfnFunctionMixinProps{
//   	CustomOutputConfiguration: &CustomOutputConfigurationProperty{
//   		Output: map[string]*string{
//   			"outputKey": jsii.String("output"),
//   		},
//   		Runtime: jsii.String("runtime"),
//   	},
//   	Description: jsii.String("description"),
//   	FunctionId: jsii.String("functionId"),
//   	FunctionType: jsii.String("functionType"),
//   	HttpRequestConfiguration: &HttpRequestConfigurationProperty{
//   		Body: jsii.String("body"),
//   		Headers: map[string]*string{
//   			"headersKey": jsii.String("headers"),
//   		},
//   		MethodType: jsii.String("methodType"),
//   		Output: map[string]*string{
//   			"outputKey": jsii.String("output"),
//   		},
//   		RequestTimeoutMilliseconds: jsii.Number(123),
//   		Runtime: jsii.String("runtime"),
//   		Url: jsii.String("url"),
//   	},
//   	SequentialExecutorConfiguration: &SequentialExecutorConfigurationProperty{
//   		FunctionList: []interface{}{
//   			&FunctionRefProperty{
//   				FunctionId: jsii.String("functionId"),
//   				RunCondition: jsii.String("runCondition"),
//   			},
//   		},
//   		Output: map[string]*string{
//   			"outputKey": jsii.String("output"),
//   		},
//   		Runtime: jsii.String("runtime"),
//   		TimeoutMilliseconds: jsii.Number(123),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html
//
type CfnFunctionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnFunctionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFunctionPropsMixin
type jsiiProxy_CfnFunctionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnFunctionPropsMixin) Props() *CfnFunctionMixinProps {
	var returns *CfnFunctionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaTailor::Function`.
func NewCfnFunctionPropsMixin(props *CfnFunctionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnFunctionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFunctionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFunctionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnFunctionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaTailor::Function`.
func NewCfnFunctionPropsMixin_Override(c CfnFunctionPropsMixin, props *CfnFunctionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnFunctionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnFunctionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFunctionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnFunctionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFunctionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnFunctionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFunctionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFunctionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


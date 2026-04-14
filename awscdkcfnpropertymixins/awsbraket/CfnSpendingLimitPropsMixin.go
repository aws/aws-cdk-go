package awsbraket

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbraket/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a spending limit for a specified quantum device.
//
// Spending limits help you control costs by setting maximum amounts that can be spent on quantum computing tasks within a specified time period.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnSpendingLimitPropsMixin := awscdkcfnpropertymixins.Aws_braket.NewCfnSpendingLimitPropsMixin(&CfnSpendingLimitMixinProps{
//   	DeviceArn: jsii.String("deviceArn"),
//   	SpendingLimit: jsii.String("spendingLimit"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimePeriod: &TimePeriodProperty{
//   		EndAt: jsii.String("endAt"),
//   		StartAt: jsii.String("startAt"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-braket-spendinglimit.html
//
type CfnSpendingLimitPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnSpendingLimitMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSpendingLimitPropsMixin
type jsiiProxy_CfnSpendingLimitPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnSpendingLimitPropsMixin) Props() *CfnSpendingLimitMixinProps {
	var returns *CfnSpendingLimitMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSpendingLimitPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Braket::SpendingLimit`.
func NewCfnSpendingLimitPropsMixin(props *CfnSpendingLimitMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnSpendingLimitPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSpendingLimitPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSpendingLimitPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_braket.CfnSpendingLimitPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Braket::SpendingLimit`.
func NewCfnSpendingLimitPropsMixin_Override(c CfnSpendingLimitPropsMixin, props *CfnSpendingLimitMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_braket.CfnSpendingLimitPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnSpendingLimitPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSpendingLimitPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_braket.CfnSpendingLimitPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSpendingLimitPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_braket.CfnSpendingLimitPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSpendingLimitPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSpendingLimitPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


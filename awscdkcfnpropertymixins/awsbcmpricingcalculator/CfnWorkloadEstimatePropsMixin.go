package awsbcmpricingcalculator

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbcmpricingcalculator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::BcmPricingCalculator::WorkloadEstimate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnWorkloadEstimatePropsMixin := awscdkcfnpropertymixins.Aws_bcmpricingcalculator.NewCfnWorkloadEstimatePropsMixin(&CfnWorkloadEstimateMixinProps{
//   	ExpiresAt: jsii.String("expiresAt"),
//   	Name: jsii.String("name"),
//   	RateType: jsii.String("rateType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-workloadestimate.html
//
type CfnWorkloadEstimatePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnWorkloadEstimateMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkloadEstimatePropsMixin
type jsiiProxy_CfnWorkloadEstimatePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnWorkloadEstimatePropsMixin) Props() *CfnWorkloadEstimateMixinProps {
	var returns *CfnWorkloadEstimateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkloadEstimatePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BcmPricingCalculator::WorkloadEstimate`.
func NewCfnWorkloadEstimatePropsMixin(props *CfnWorkloadEstimateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnWorkloadEstimatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkloadEstimatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkloadEstimatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bcmpricingcalculator.CfnWorkloadEstimatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BcmPricingCalculator::WorkloadEstimate`.
func NewCfnWorkloadEstimatePropsMixin_Override(c CfnWorkloadEstimatePropsMixin, props *CfnWorkloadEstimateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bcmpricingcalculator.CfnWorkloadEstimatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnWorkloadEstimatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkloadEstimatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bcmpricingcalculator.CfnWorkloadEstimatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkloadEstimatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bcmpricingcalculator.CfnWorkloadEstimatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkloadEstimatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnWorkloadEstimatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


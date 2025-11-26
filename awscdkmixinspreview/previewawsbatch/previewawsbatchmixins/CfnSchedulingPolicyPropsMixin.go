package previewawsbatchmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbatch/previewawsbatchmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Batch::SchedulingPolicy` resource specifies the parameters for an AWS Batch scheduling policy.
//
// For more information, see [Scheduling Policies](https://docs.aws.amazon.com/batch/latest/userguide/scheduling_policies.html) in the ** .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSchedulingPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnSchedulingPolicyPropsMixin(&CfnSchedulingPolicyMixinProps{
//   	FairsharePolicy: &FairsharePolicyProperty{
//   		ComputeReservation: jsii.Number(123),
//   		ShareDecaySeconds: jsii.Number(123),
//   		ShareDistribution: []interface{}{
//   			&ShareAttributesProperty{
//   				ShareIdentifier: jsii.String("shareIdentifier"),
//   				WeightFactor: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-schedulingpolicy.html
//
type CfnSchedulingPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSchedulingPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSchedulingPolicyPropsMixin
type jsiiProxy_CfnSchedulingPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSchedulingPolicyPropsMixin) Props() *CfnSchedulingPolicyMixinProps {
	var returns *CfnSchedulingPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedulingPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Batch::SchedulingPolicy`.
func NewCfnSchedulingPolicyPropsMixin(props *CfnSchedulingPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSchedulingPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSchedulingPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSchedulingPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnSchedulingPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Batch::SchedulingPolicy`.
func NewCfnSchedulingPolicyPropsMixin_Override(c CfnSchedulingPolicyPropsMixin, props *CfnSchedulingPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnSchedulingPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSchedulingPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSchedulingPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnSchedulingPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSchedulingPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnSchedulingPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSchedulingPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSchedulingPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


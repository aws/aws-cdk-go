package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Batch::QuotaShare.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnQuotaSharePropsMixin := awscdkcfnpropertymixins.Aws_batch.NewCfnQuotaSharePropsMixin(&CfnQuotaShareMixinProps{
//   	CapacityLimits: []interface{}{
//   		&QuotaShareCapacityLimitProperty{
//   			CapacityUnit: jsii.String("capacityUnit"),
//   			MaxCapacity: jsii.Number(123),
//   		},
//   	},
//   	JobQueue: jsii.String("jobQueue"),
//   	PreemptionConfiguration: &QuotaSharePreemptionConfigurationProperty{
//   		InSharePreemption: jsii.String("inSharePreemption"),
//   	},
//   	QuotaShareName: jsii.String("quotaShareName"),
//   	ResourceSharingConfiguration: &QuotaShareResourceSharingConfigurationProperty{
//   		BorrowLimit: jsii.Number(123),
//   		Strategy: jsii.String("strategy"),
//   	},
//   	State: jsii.String("state"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-quotashare.html
//
type CfnQuotaSharePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnQuotaShareMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnQuotaSharePropsMixin
type jsiiProxy_CfnQuotaSharePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnQuotaSharePropsMixin) Props() *CfnQuotaShareMixinProps {
	var returns *CfnQuotaShareMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQuotaSharePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Batch::QuotaShare`.
func NewCfnQuotaSharePropsMixin(props *CfnQuotaShareMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnQuotaSharePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnQuotaSharePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnQuotaSharePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnQuotaSharePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Batch::QuotaShare`.
func NewCfnQuotaSharePropsMixin_Override(c CfnQuotaSharePropsMixin, props *CfnQuotaShareMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnQuotaSharePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnQuotaSharePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnQuotaSharePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnQuotaSharePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnQuotaSharePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_batch.CfnQuotaSharePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnQuotaSharePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnQuotaSharePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


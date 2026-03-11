package awscustomerprofiles

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies the rules to perform an action based on customer ingested data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnEventTriggerPropsMixin := awscdkcfnpropertymixins.Aws_customerprofiles.NewCfnEventTriggerPropsMixin(&CfnEventTriggerMixinProps{
//   	Description: jsii.String("description"),
//   	DomainName: jsii.String("domainName"),
//   	EventTriggerConditions: []interface{}{
//   		&EventTriggerConditionProperty{
//   			EventTriggerDimensions: []interface{}{
//   				&EventTriggerDimensionProperty{
//   					ObjectAttributes: []interface{}{
//   						&ObjectAttributeProperty{
//   							ComparisonOperator: jsii.String("comparisonOperator"),
//   							FieldName: jsii.String("fieldName"),
//   							Source: jsii.String("source"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			LogicalOperator: jsii.String("logicalOperator"),
//   		},
//   	},
//   	EventTriggerLimits: &EventTriggerLimitsProperty{
//   		EventExpiration: jsii.Number(123),
//   		Periods: []interface{}{
//   			&PeriodProperty{
//   				MaxInvocationsPerProfile: jsii.Number(123),
//   				Unit: jsii.String("unit"),
//   				Unlimited: jsii.Boolean(false),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	EventTriggerName: jsii.String("eventTriggerName"),
//   	ObjectTypeName: jsii.String("objectTypeName"),
//   	SegmentFilter: jsii.String("segmentFilter"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html
//
type CfnEventTriggerPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnEventTriggerMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEventTriggerPropsMixin
type jsiiProxy_CfnEventTriggerPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnEventTriggerPropsMixin) Props() *CfnEventTriggerMixinProps {
	var returns *CfnEventTriggerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventTriggerPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CustomerProfiles::EventTrigger`.
func NewCfnEventTriggerPropsMixin(props *CfnEventTriggerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnEventTriggerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEventTriggerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEventTriggerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnEventTriggerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CustomerProfiles::EventTrigger`.
func NewCfnEventTriggerPropsMixin_Override(c CfnEventTriggerPropsMixin, props *CfnEventTriggerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnEventTriggerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnEventTriggerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventTriggerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnEventTriggerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventTriggerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_customerprofiles.CfnEventTriggerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEventTriggerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEventTriggerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


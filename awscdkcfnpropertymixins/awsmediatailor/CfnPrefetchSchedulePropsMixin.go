package awsmediatailor

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::MediaTailor::PrefetchSchedule Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnPrefetchSchedulePropsMixin := awscdkcfnpropertymixins.Aws_mediatailor.NewCfnPrefetchSchedulePropsMixin(&CfnPrefetchScheduleMixinProps{
//   	Consumption: &PrefetchConsumptionProperty{
//   		AvailMatchingCriteria: []interface{}{
//   			&AvailMatchingCriteriaProperty{
//   				DynamicVariable: jsii.String("dynamicVariable"),
//   				Operator: jsii.String("operator"),
//   			},
//   		},
//   		EndTime: jsii.String("endTime"),
//   		StartTime: jsii.String("startTime"),
//   	},
//   	Name: jsii.String("name"),
//   	PlaybackConfigurationName: jsii.String("playbackConfigurationName"),
//   	RecurringPrefetchConfiguration: &RecurringPrefetchConfigurationProperty{
//   		EndTime: jsii.String("endTime"),
//   		RecurringConsumption: &RecurringConsumptionProperty{
//   			AvailMatchingCriteria: []interface{}{
//   				&AvailMatchingCriteriaProperty{
//   					DynamicVariable: jsii.String("dynamicVariable"),
//   					Operator: jsii.String("operator"),
//   				},
//   			},
//   			RetrievedAdExpirationSeconds: jsii.Number(123),
//   		},
//   		RecurringRetrieval: &RecurringRetrievalProperty{
//   			DelayAfterAvailEndSeconds: jsii.Number(123),
//   			DynamicVariables: map[string]*string{
//   				"dynamicVariablesKey": jsii.String("dynamicVariables"),
//   			},
//   			TrafficShapingRetrievalWindow: &TrafficShapingRetrievalWindowProperty{
//   				RetrievalWindowDurationSeconds: jsii.Number(123),
//   			},
//   			TrafficShapingTpsConfiguration: &TrafficShapingTpsConfigurationProperty{
//   				PeakConcurrentUsers: jsii.Number(123),
//   				PeakTps: jsii.Number(123),
//   			},
//   			TrafficShapingType: jsii.String("trafficShapingType"),
//   		},
//   		StartTime: jsii.String("startTime"),
//   	},
//   	Retrieval: &PrefetchRetrievalProperty{
//   		DynamicVariables: map[string]*string{
//   			"dynamicVariablesKey": jsii.String("dynamicVariables"),
//   		},
//   		EndTime: jsii.String("endTime"),
//   		StartTime: jsii.String("startTime"),
//   		TrafficShapingRetrievalWindow: &TrafficShapingRetrievalWindowProperty{
//   			RetrievalWindowDurationSeconds: jsii.Number(123),
//   		},
//   		TrafficShapingTpsConfiguration: &TrafficShapingTpsConfigurationProperty{
//   			PeakConcurrentUsers: jsii.Number(123),
//   			PeakTps: jsii.Number(123),
//   		},
//   		TrafficShapingType: jsii.String("trafficShapingType"),
//   	},
//   	ScheduleType: jsii.String("scheduleType"),
//   	StreamId: jsii.String("streamId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-prefetchschedule.html
//
type CfnPrefetchSchedulePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPrefetchScheduleMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPrefetchSchedulePropsMixin
type jsiiProxy_CfnPrefetchSchedulePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnPrefetchSchedulePropsMixin) Props() *CfnPrefetchScheduleMixinProps {
	var returns *CfnPrefetchScheduleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrefetchSchedulePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaTailor::PrefetchSchedule`.
func NewCfnPrefetchSchedulePropsMixin(props *CfnPrefetchScheduleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPrefetchSchedulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPrefetchSchedulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPrefetchSchedulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPrefetchSchedulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaTailor::PrefetchSchedule`.
func NewCfnPrefetchSchedulePropsMixin_Override(c CfnPrefetchSchedulePropsMixin, props *CfnPrefetchScheduleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPrefetchSchedulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPrefetchSchedulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPrefetchSchedulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPrefetchSchedulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPrefetchSchedulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnPrefetchSchedulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPrefetchSchedulePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPrefetchSchedulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


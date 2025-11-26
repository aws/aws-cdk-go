package previewawsevidentlymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsevidently/previewawsevidentlymixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates or updates a *launch* of a given feature.
//
// Before you create a launch, you must create the feature to use for the launch.
//
// You can use a launch to safely validate new features by serving them to a specified percentage of your users while you roll out the feature. You can monitor the performance of the new feature to help you decide when to ramp up traffic to more users. This helps you reduce risk and identify unintended consequences before you fully launch the feature.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLaunchPropsMixin := awscdkmixinspreview.Mixins.NewCfnLaunchPropsMixin(&CfnLaunchMixinProps{
//   	Description: jsii.String("description"),
//   	ExecutionStatus: &ExecutionStatusObjectProperty{
//   		DesiredState: jsii.String("desiredState"),
//   		Reason: jsii.String("reason"),
//   		Status: jsii.String("status"),
//   	},
//   	Groups: []interface{}{
//   		&LaunchGroupObjectProperty{
//   			Description: jsii.String("description"),
//   			Feature: jsii.String("feature"),
//   			GroupName: jsii.String("groupName"),
//   			Variation: jsii.String("variation"),
//   		},
//   	},
//   	MetricMonitors: []interface{}{
//   		&MetricDefinitionObjectProperty{
//   			EntityIdKey: jsii.String("entityIdKey"),
//   			EventPattern: jsii.String("eventPattern"),
//   			MetricName: jsii.String("metricName"),
//   			UnitLabel: jsii.String("unitLabel"),
//   			ValueKey: jsii.String("valueKey"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Project: jsii.String("project"),
//   	RandomizationSalt: jsii.String("randomizationSalt"),
//   	ScheduledSplitsConfig: []interface{}{
//   		&StepConfigProperty{
//   			GroupWeights: []interface{}{
//   				&GroupToWeightProperty{
//   					GroupName: jsii.String("groupName"),
//   					SplitWeight: jsii.Number(123),
//   				},
//   			},
//   			SegmentOverrides: []interface{}{
//   				&SegmentOverrideProperty{
//   					EvaluationOrder: jsii.Number(123),
//   					Segment: jsii.String("segment"),
//   					Weights: []interface{}{
//   						&GroupToWeightProperty{
//   							GroupName: jsii.String("groupName"),
//   							SplitWeight: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   			StartTime: jsii.String("startTime"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html
//
type CfnLaunchPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLaunchMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLaunchPropsMixin
type jsiiProxy_CfnLaunchPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLaunchPropsMixin) Props() *CfnLaunchMixinProps {
	var returns *CfnLaunchMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Evidently::Launch`.
func NewCfnLaunchPropsMixin(props *CfnLaunchMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLaunchPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLaunchPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLaunchPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnLaunchPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Evidently::Launch`.
func NewCfnLaunchPropsMixin_Override(c CfnLaunchPropsMixin, props *CfnLaunchMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnLaunchPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLaunchPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLaunchPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnLaunchPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLaunchPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnLaunchPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLaunchPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLaunchPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


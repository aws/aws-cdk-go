package mixinsawsevidently

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsevidently/mixinsawsevidently/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates or updates an Evidently *experiment* .
//
// Before you create an experiment, you must create the feature to use for the experiment.
//
// An experiment helps you make feature design decisions based on evidence and data. An experiment can test as many as five variations at once. Evidently collects experiment data and analyzes it by statistical methods, and provides clear recommendations about which variations perform better.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnExperimentPropsMixin := awscdkmixinspreview.Mixins.NewCfnExperimentPropsMixin(&CfnExperimentMixinProps{
//   	Description: jsii.String("description"),
//   	MetricGoals: []interface{}{
//   		&MetricGoalObjectProperty{
//   			DesiredChange: jsii.String("desiredChange"),
//   			EntityIdKey: jsii.String("entityIdKey"),
//   			EventPattern: jsii.String("eventPattern"),
//   			MetricName: jsii.String("metricName"),
//   			UnitLabel: jsii.String("unitLabel"),
//   			ValueKey: jsii.String("valueKey"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	OnlineAbConfig: &OnlineAbConfigObjectProperty{
//   		ControlTreatmentName: jsii.String("controlTreatmentName"),
//   		TreatmentWeights: []interface{}{
//   			&TreatmentToWeightProperty{
//   				SplitWeight: jsii.Number(123),
//   				Treatment: jsii.String("treatment"),
//   			},
//   		},
//   	},
//   	Project: jsii.String("project"),
//   	RandomizationSalt: jsii.String("randomizationSalt"),
//   	RemoveSegment: jsii.Boolean(false),
//   	RunningStatus: &RunningStatusObjectProperty{
//   		AnalysisCompleteTime: jsii.String("analysisCompleteTime"),
//   		DesiredState: jsii.String("desiredState"),
//   		Reason: jsii.String("reason"),
//   		Status: jsii.String("status"),
//   	},
//   	SamplingRate: jsii.Number(123),
//   	Segment: jsii.String("segment"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Treatments: []interface{}{
//   		&TreatmentObjectProperty{
//   			Description: jsii.String("description"),
//   			Feature: jsii.String("feature"),
//   			TreatmentName: jsii.String("treatmentName"),
//   			Variation: jsii.String("variation"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-experiment.html
//
type CfnExperimentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnExperimentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnExperimentPropsMixin
type jsiiProxy_CfnExperimentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnExperimentPropsMixin) Props() *CfnExperimentMixinProps {
	var returns *CfnExperimentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Evidently::Experiment`.
func NewCfnExperimentPropsMixin(props *CfnExperimentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnExperimentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnExperimentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnExperimentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnExperimentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Evidently::Experiment`.
func NewCfnExperimentPropsMixin_Override(c CfnExperimentPropsMixin, props *CfnExperimentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnExperimentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnExperimentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnExperimentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnExperimentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnExperimentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnExperimentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnExperimentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnExperimentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


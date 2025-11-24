package mixinsawsaps

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsaps/mixinsawsaps/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Anomaly detection uses the Random Cut Forest algorithm for time-series analysis.
//
// The anomaly detector analyzes Amazon Managed Service for Prometheus metrics to identify unusual patterns and behaviors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAnomalyDetectorPropsMixin := awscdkmixinspreview.Mixins.NewCfnAnomalyDetectorPropsMixin(&CfnAnomalyDetectorMixinProps{
//   	Alias: jsii.String("alias"),
//   	Configuration: &AnomalyDetectorConfigurationProperty{
//   		RandomCutForest: &RandomCutForestConfigurationProperty{
//   			IgnoreNearExpectedFromAbove: &IgnoreNearExpectedProperty{
//   				Amount: jsii.Number(123),
//   				Ratio: jsii.Number(123),
//   			},
//   			IgnoreNearExpectedFromBelow: &IgnoreNearExpectedProperty{
//   				Amount: jsii.Number(123),
//   				Ratio: jsii.Number(123),
//   			},
//   			Query: jsii.String("query"),
//   			SampleSize: jsii.Number(123),
//   			ShingleSize: jsii.Number(123),
//   		},
//   	},
//   	EvaluationIntervalInSeconds: jsii.Number(123),
//   	Labels: []interface{}{
//   		&LabelProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MissingDataAction: &MissingDataActionProperty{
//   		MarkAsAnomaly: jsii.Boolean(false),
//   		Skip: jsii.Boolean(false),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Workspace: jsii.String("workspace"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-anomalydetector.html
//
type CfnAnomalyDetectorPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAnomalyDetectorMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAnomalyDetectorPropsMixin
type jsiiProxy_CfnAnomalyDetectorPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAnomalyDetectorPropsMixin) Props() *CfnAnomalyDetectorMixinProps {
	var returns *CfnAnomalyDetectorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetectorPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::APS::AnomalyDetector`.
func NewCfnAnomalyDetectorPropsMixin(props *CfnAnomalyDetectorMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAnomalyDetectorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAnomalyDetectorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAnomalyDetectorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnAnomalyDetectorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::APS::AnomalyDetector`.
func NewCfnAnomalyDetectorPropsMixin_Override(c CfnAnomalyDetectorPropsMixin, props *CfnAnomalyDetectorMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnAnomalyDetectorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAnomalyDetectorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAnomalyDetectorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnAnomalyDetectorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAnomalyDetectorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnAnomalyDetectorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAnomalyDetectorPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAnomalyDetectorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


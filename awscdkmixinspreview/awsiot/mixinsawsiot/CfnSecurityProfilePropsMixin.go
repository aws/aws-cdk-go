package mixinsawsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsiot/mixinsawsiot/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the `AWS::IoT::SecurityProfile` resource to create a Device Defender security profile.
//
// For API reference, see [CreateSecurityProfile](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateSecurityProfile.html) and for general information, see [Detect](https://docs.aws.amazon.com/iot/latest/developerguide/device-defender-detect.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSecurityProfilePropsMixin := awscdkmixinspreview.Mixins.NewCfnSecurityProfilePropsMixin(&CfnSecurityProfileMixinProps{
//   	AdditionalMetricsToRetainV2: []interface{}{
//   		&MetricToRetainProperty{
//   			ExportMetric: jsii.Boolean(false),
//   			Metric: jsii.String("metric"),
//   			MetricDimension: &MetricDimensionProperty{
//   				DimensionName: jsii.String("dimensionName"),
//   				Operator: jsii.String("operator"),
//   			},
//   		},
//   	},
//   	AlertTargets: map[string]interface{}{
//   		"alertTargetsKey": &AlertTargetProperty{
//   			"alertTargetArn": jsii.String("alertTargetArn"),
//   			"roleArn": jsii.String("roleArn"),
//   		},
//   	},
//   	Behaviors: []interface{}{
//   		&BehaviorProperty{
//   			Criteria: &BehaviorCriteriaProperty{
//   				ComparisonOperator: jsii.String("comparisonOperator"),
//   				ConsecutiveDatapointsToAlarm: jsii.Number(123),
//   				ConsecutiveDatapointsToClear: jsii.Number(123),
//   				DurationSeconds: jsii.Number(123),
//   				MlDetectionConfig: &MachineLearningDetectionConfigProperty{
//   					ConfidenceLevel: jsii.String("confidenceLevel"),
//   				},
//   				StatisticalThreshold: &StatisticalThresholdProperty{
//   					Statistic: jsii.String("statistic"),
//   				},
//   				Value: &MetricValueProperty{
//   					Cidrs: []*string{
//   						jsii.String("cidrs"),
//   					},
//   					Count: jsii.String("count"),
//   					Number: jsii.Number(123),
//   					Numbers: []interface{}{
//   						jsii.Number(123),
//   					},
//   					Ports: []interface{}{
//   						jsii.Number(123),
//   					},
//   					Strings: []*string{
//   						jsii.String("strings"),
//   					},
//   				},
//   			},
//   			ExportMetric: jsii.Boolean(false),
//   			Metric: jsii.String("metric"),
//   			MetricDimension: &MetricDimensionProperty{
//   				DimensionName: jsii.String("dimensionName"),
//   				Operator: jsii.String("operator"),
//   			},
//   			Name: jsii.String("name"),
//   			SuppressAlerts: jsii.Boolean(false),
//   		},
//   	},
//   	MetricsExportConfig: &MetricsExportConfigProperty{
//   		MqttTopic: jsii.String("mqttTopic"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	SecurityProfileDescription: jsii.String("securityProfileDescription"),
//   	SecurityProfileName: jsii.String("securityProfileName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetArns: []*string{
//   		jsii.String("targetArns"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html
//
type CfnSecurityProfilePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSecurityProfileMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSecurityProfilePropsMixin
type jsiiProxy_CfnSecurityProfilePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSecurityProfilePropsMixin) Props() *CfnSecurityProfileMixinProps {
	var returns *CfnSecurityProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfilePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::SecurityProfile`.
func NewCfnSecurityProfilePropsMixin(props *CfnSecurityProfileMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSecurityProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSecurityProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSecurityProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnSecurityProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::SecurityProfile`.
func NewCfnSecurityProfilePropsMixin_Override(c CfnSecurityProfilePropsMixin, props *CfnSecurityProfileMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnSecurityProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSecurityProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSecurityProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnSecurityProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecurityProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnSecurityProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSecurityProfilePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSecurityProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::CloudWatch::Alert.
//
// An alert runs a SQL or PromQL query in a CloudWatch Omni space on a fixed interval, compares the result against warning and critical thresholds, and sends notifications to Amazon SNS topics or integrations when its state changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAlertPropsMixin := awscdkcfnpropertymixins.Aws_cloudwatch.NewCfnAlertPropsMixin(&CfnAlertMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	NotificationRules: []interface{}{
//   		&NotificationRuleProperty{
//   			Target: &TargetProperty{
//   				Arn: jsii.String("arn"),
//   				Metadata: map[string]*string{
//   					"metadataKey": jsii.String("metadata"),
//   				},
//   				Type: jsii.String("type"),
//   			},
//   			Trigger: &TriggerProperty{
//   				StateValues: []*string{
//   					jsii.String("stateValues"),
//   				},
//   			},
//   		},
//   	},
//   	NotificationsEnabled: jsii.Boolean(false),
//   	ProfileId: jsii.String("profileId"),
//   	Rule: &RuleProperty{
//   		TelemetryRule: &TelemetryRuleProperty{
//   			Condition: &ConditionProperty{
//   				Comparator: jsii.String("comparator"),
//   				CriticalThreshold: jsii.Number(123),
//   				ThresholdField: jsii.String("thresholdField"),
//   				ThresholdMode: jsii.String("thresholdMode"),
//   				WarningThreshold: jsii.Number(123),
//   			},
//   			Evaluation: &EvaluationProperty{
//   				IntervalSeconds: jsii.Number(123),
//   				PendingDurationSeconds: jsii.Number(123),
//   				RecoveryDurationSeconds: jsii.Number(123),
//   			},
//   			NoData: &NoDataProperty{
//   				TreatAs: jsii.String("treatAs"),
//   			},
//   			Query: &QueryProperty{
//   				Expression: jsii.String("expression"),
//   				Language: jsii.String("language"),
//   			},
//   		},
//   	},
//   	SpaceId: jsii.String("spaceId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alert.html
//
type CfnAlertPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAlertMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAlertPropsMixin
type jsiiProxy_CfnAlertPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAlertPropsMixin) Props() *CfnAlertMixinProps {
	var returns *CfnAlertMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlertPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudWatch::Alert`.
func NewCfnAlertPropsMixin(props *CfnAlertMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAlertPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAlertPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAlertPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudWatch::Alert`.
func NewCfnAlertPropsMixin_Override(c CfnAlertPropsMixin, props *CfnAlertMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAlertPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAlertPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAlertPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAlertPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAlertPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


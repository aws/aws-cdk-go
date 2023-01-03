package awsstepfunctionstasks


// A scale-in or scale-out rule that defines scaling activity, including the CloudWatch metric alarm that triggers activity, how EC2 instances are added or removed, and the periodicity of adjustments.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingRuleProperty := &scalingRuleProperty{
//   	action: &scalingActionProperty{
//   		simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   			scalingAdjustment: jsii.Number(123),
//
//   			// the properties below are optional
//   			adjustmentType: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.scalingAdjustmentType_CHANGE_IN_CAPACITY,
//   			coolDown: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		market: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.instanceMarket_ON_DEMAND,
//   	},
//   	name: jsii.String("name"),
//   	trigger: &scalingTriggerProperty{
//   		cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   			comparisonOperator: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmComparisonOperator_GREATER_THAN_OR_EQUAL,
//   			metricName: jsii.String("metricName"),
//   			period: cdk.duration.minutes(jsii.Number(30)),
//
//   			// the properties below are optional
//   			dimensions: []metricDimensionProperty{
//   				&metricDimensionProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			evaluationPeriods: jsii.Number(123),
//   			namespace: jsii.String("namespace"),
//   			statistic: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmStatistic_SAMPLE_COUNT,
//   			threshold: jsii.Number(123),
//   			unit: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmUnit_NONE,
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ScalingRule.html
//
type EmrCreateCluster_ScalingRuleProperty struct {
	// The conditions that trigger an automatic scaling activity.
	Action *EmrCreateCluster_ScalingActionProperty `field:"required" json:"action" yaml:"action"`
	// The name used to identify an automatic scaling rule.
	//
	// Rule names must be unique within a scaling policy.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The CloudWatch alarm definition that determines when automatic scaling activity is triggered.
	Trigger *EmrCreateCluster_ScalingTriggerProperty `field:"required" json:"trigger" yaml:"trigger"`
	// A friendly, more verbose description of the automatic scaling rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
}


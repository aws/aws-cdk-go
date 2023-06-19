package awsstepfunctionstasks


// A scale-in or scale-out rule that defines scaling activity, including the CloudWatch metric alarm that triggers activity, how EC2 instances are added or removed, and the periodicity of adjustments.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   scalingRuleProperty := &ScalingRuleProperty{
//   	Action: &ScalingActionProperty{
//   		SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   			ScalingAdjustment: jsii.Number(123),
//
//   			// the properties below are optional
//   			AdjustmentType: awscdk.Aws_stepfunctions_tasks.EmrCreateCluster.ScalingAdjustmentType_CHANGE_IN_CAPACITY,
//   			CoolDown: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		Market: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.InstanceMarket_ON_DEMAND,
//   	},
//   	Name: jsii.String("name"),
//   	Trigger: &ScalingTriggerProperty{
//   		CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   			ComparisonOperator: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmComparisonOperator_GREATER_THAN_OR_EQUAL,
//   			MetricName: jsii.String("metricName"),
//   			Period: duration,
//
//   			// the properties below are optional
//   			Dimensions: []metricDimensionProperty{
//   				&metricDimensionProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			EvaluationPeriods: jsii.Number(123),
//   			Namespace: jsii.String("namespace"),
//   			Statistic: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmStatistic_SAMPLE_COUNT,
//   			Threshold: jsii.Number(123),
//   			Unit: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmUnit_NONE,
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ScalingRule.html
//
// Experimental.
type EmrCreateCluster_ScalingRuleProperty struct {
	// The conditions that trigger an automatic scaling activity.
	// Experimental.
	Action *EmrCreateCluster_ScalingActionProperty `field:"required" json:"action" yaml:"action"`
	// The name used to identify an automatic scaling rule.
	//
	// Rule names must be unique within a scaling policy.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The CloudWatch alarm definition that determines when automatic scaling activity is triggered.
	// Experimental.
	Trigger *EmrCreateCluster_ScalingTriggerProperty `field:"required" json:"trigger" yaml:"trigger"`
	// A friendly, more verbose description of the automatic scaling rule.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}


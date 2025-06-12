package awsstepfunctionstasks


// An automatic scaling policy for a core instance group or task instance group in an Amazon EMR cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingPolicyProperty := &AutoScalingPolicyProperty{
//   	Constraints: &ScalingConstraintsProperty{
//   		MaxCapacity: jsii.Number(123),
//   		MinCapacity: jsii.Number(123),
//   	},
//   	Rules: []scalingRuleProperty{
//   		&scalingRuleProperty{
//   			Action: &ScalingActionProperty{
//   				SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   					ScalingAdjustment: jsii.Number(123),
//
//   					// the properties below are optional
//   					AdjustmentType: awscdk.Aws_stepfunctions_tasks.EmrCreateCluster.ScalingAdjustmentType_CHANGE_IN_CAPACITY,
//   					CoolDown: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				Market: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.InstanceMarket_ON_DEMAND,
//   			},
//   			Name: jsii.String("name"),
//   			Trigger: &ScalingTriggerProperty{
//   				CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   					ComparisonOperator: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmComparisonOperator_GREATER_THAN_OR_EQUAL,
//   					MetricName: jsii.String("metricName"),
//   					Period: cdk.Duration_Minutes(jsii.Number(30)),
//
//   					// the properties below are optional
//   					Dimensions: []metricDimensionProperty{
//   						&metricDimensionProperty{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					EvaluationPeriods: jsii.Number(123),
//   					Namespace: jsii.String("namespace"),
//   					Statistic: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmStatistic_SAMPLE_COUNT,
//   					Threshold: jsii.Number(123),
//   					Unit: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmUnit_NONE,
//   				},
//   			},
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_AutoScalingPolicy.html
//
type EmrCreateCluster_AutoScalingPolicyProperty struct {
	// The upper and lower EC2 instance limits for an automatic scaling policy.
	//
	// Automatic scaling activity will not cause an instance
	// group to grow above or below these limits.
	Constraints *EmrCreateCluster_ScalingConstraintsProperty `field:"required" json:"constraints" yaml:"constraints"`
	// The scale-in and scale-out rules that comprise the automatic scaling policy.
	Rules *[]*EmrCreateCluster_ScalingRuleProperty `field:"required" json:"rules" yaml:"rules"`
}


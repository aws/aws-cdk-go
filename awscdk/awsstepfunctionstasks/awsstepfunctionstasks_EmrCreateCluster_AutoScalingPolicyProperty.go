package awsstepfunctionstasks


// An automatic scaling policy for a core instance group or task instance group in an Amazon EMR cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingPolicyProperty := &autoScalingPolicyProperty{
//   	constraints: &scalingConstraintsProperty{
//   		maxCapacity: jsii.Number(123),
//   		minCapacity: jsii.Number(123),
//   	},
//   	rules: []scalingRuleProperty{
//   		&scalingRuleProperty{
//   			action: &scalingActionProperty{
//   				simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   					scalingAdjustment: jsii.Number(123),
//
//   					// the properties below are optional
//   					adjustmentType: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.scalingAdjustmentType_CHANGE_IN_CAPACITY,
//   					coolDown: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				market: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.instanceMarket_ON_DEMAND,
//   			},
//   			name: jsii.String("name"),
//   			trigger: &scalingTriggerProperty{
//   				cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   					comparisonOperator: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmComparisonOperator_GREATER_THAN_OR_EQUAL,
//   					metricName: jsii.String("metricName"),
//   					period: cdk.duration.minutes(jsii.Number(30)),
//
//   					// the properties below are optional
//   					dimensions: []metricDimensionProperty{
//   						&metricDimensionProperty{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					evaluationPeriods: jsii.Number(123),
//   					namespace: jsii.String("namespace"),
//   					statistic: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmStatistic_SAMPLE_COUNT,
//   					threshold: jsii.Number(123),
//   					unit: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmUnit_NONE,
//   				},
//   			},
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
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


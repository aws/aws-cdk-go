package awsstepfunctionstasks


// The conditions that trigger an automatic scaling activity and the definition of a CloudWatch metric alarm.
//
// When the defined alarm conditions are met along with other trigger parameters, scaling activity begins.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   scalingTriggerProperty := &ScalingTriggerProperty{
//   	CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   		ComparisonOperator: awscdk.Aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmComparisonOperator_GREATER_THAN_OR_EQUAL,
//   		MetricName: jsii.String("metricName"),
//   		Period: duration,
//
//   		// the properties below are optional
//   		Dimensions: []metricDimensionProperty{
//   			&metricDimensionProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		EvaluationPeriods: jsii.Number(123),
//   		Namespace: jsii.String("namespace"),
//   		Statistic: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmStatistic_SAMPLE_COUNT,
//   		Threshold: jsii.Number(123),
//   		Unit: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmUnit_NONE,
//   	},
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ScalingTrigger.html
//
// Experimental.
type EmrCreateCluster_ScalingTriggerProperty struct {
	// The definition of a CloudWatch metric alarm.
	//
	// When the defined alarm conditions are met along with other trigger parameters,
	// scaling activity begins.
	// Experimental.
	CloudWatchAlarmDefinition *EmrCreateCluster_CloudWatchAlarmDefinitionProperty `field:"required" json:"cloudWatchAlarmDefinition" yaml:"cloudWatchAlarmDefinition"`
}


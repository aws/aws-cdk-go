package awsstepfunctionstasks


// The conditions that trigger an automatic scaling activity and the definition of a CloudWatch metric alarm.
//
// When the defined alarm conditions are met along with other trigger parameters, scaling activity begins.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingTriggerProperty := &scalingTriggerProperty{
//   	cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   		comparisonOperator: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.cloudWatchAlarmComparisonOperator_GREATER_THAN_OR_EQUAL,
//   		metricName: jsii.String("metricName"),
//   		period: cdk.duration.minutes(jsii.Number(30)),
//
//   		// the properties below are optional
//   		dimensions: []metricDimensionProperty{
//   			&metricDimensionProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		evaluationPeriods: jsii.Number(123),
//   		namespace: jsii.String("namespace"),
//   		statistic: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmStatistic_SAMPLE_COUNT,
//   		threshold: jsii.Number(123),
//   		unit: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmUnit_NONE,
//   	},
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ScalingTrigger.html
//
type EmrCreateCluster_ScalingTriggerProperty struct {
	// The definition of a CloudWatch metric alarm.
	//
	// When the defined alarm conditions are met along with other trigger parameters,
	// scaling activity begins.
	CloudWatchAlarmDefinition *EmrCreateCluster_CloudWatchAlarmDefinitionProperty `field:"required" json:"cloudWatchAlarmDefinition" yaml:"cloudWatchAlarmDefinition"`
}


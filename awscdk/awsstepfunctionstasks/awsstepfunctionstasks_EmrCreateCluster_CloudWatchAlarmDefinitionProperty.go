package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// The definition of a CloudWatch metric alarm, which determines when an automatic scaling activity is triggered.
//
// When the defined alarm conditions
// are satisfied, scaling activity begins.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   cloudWatchAlarmDefinitionProperty := &cloudWatchAlarmDefinitionProperty{
//   	comparisonOperator: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.cloudWatchAlarmComparisonOperator_GREATER_THAN_OR_EQUAL,
//   	metricName: jsii.String("metricName"),
//   	period: duration,
//
//   	// the properties below are optional
//   	dimensions: []metricDimensionProperty{
//   		&metricDimensionProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	evaluationPeriods: jsii.Number(123),
//   	namespace: jsii.String("namespace"),
//   	statistic: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmStatistic_SAMPLE_COUNT,
//   	threshold: jsii.Number(123),
//   	unit: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmUnit_NONE,
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_CloudWatchAlarmDefinition.html
//
// Experimental.
type EmrCreateCluster_CloudWatchAlarmDefinitionProperty struct {
	// Determines how the metric specified by MetricName is compared to the value specified by Threshold.
	// Experimental.
	ComparisonOperator EmrCreateCluster_CloudWatchAlarmComparisonOperator `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The name of the CloudWatch metric that is watched to determine an alarm condition.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The period, in seconds, over which the statistic is applied.
	//
	// EMR CloudWatch metrics are emitted every five minutes (300 seconds), so if
	// an EMR CloudWatch metric is specified, specify 300.
	// Experimental.
	Period awscdk.Duration `field:"required" json:"period" yaml:"period"`
	// A CloudWatch metric dimension.
	// Experimental.
	Dimensions *[]*EmrCreateCluster_MetricDimensionProperty `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The number of periods, in five-minute increments, during which the alarm condition must exist before the alarm triggers automatic scaling activity.
	// Experimental.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The namespace for the CloudWatch metric.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The statistic to apply to the metric associated with the alarm.
	// Experimental.
	Statistic EmrCreateCluster_CloudWatchAlarmStatistic `field:"optional" json:"statistic" yaml:"statistic"`
	// The value against which the specified statistic is compared.
	// Experimental.
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// The unit of measure associated with the CloudWatch metric being watched.
	//
	// The value specified for Unit must correspond to the units
	// specified in the CloudWatch metric.
	// Experimental.
	Unit EmrCreateCluster_CloudWatchAlarmUnit `field:"optional" json:"unit" yaml:"unit"`
}


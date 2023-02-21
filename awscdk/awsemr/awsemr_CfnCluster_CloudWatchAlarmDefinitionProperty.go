package awsemr


// `CloudWatchAlarmDefinition` is a subproperty of the `ScalingTrigger` property, which determines when to trigger an automatic scaling activity.
//
// Scaling activity begins when you satisfy the defined alarm conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchAlarmDefinitionProperty := &cloudWatchAlarmDefinitionProperty{
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	metricName: jsii.String("metricName"),
//   	period: jsii.Number(123),
//   	threshold: jsii.Number(123),
//
//   	// the properties below are optional
//   	dimensions: []interface{}{
//   		&metricDimensionProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	evaluationPeriods: jsii.Number(123),
//   	namespace: jsii.String("namespace"),
//   	statistic: jsii.String("statistic"),
//   	unit: jsii.String("unit"),
//   }
//
type CfnCluster_CloudWatchAlarmDefinitionProperty struct {
	// Determines how the metric specified by `MetricName` is compared to the value specified by `Threshold` .
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The name of the CloudWatch metric that is watched to determine an alarm condition.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The period, in seconds, over which the statistic is applied.
	//
	// EMR CloudWatch metrics are emitted every five minutes (300 seconds), so if an EMR CloudWatch metric is specified, specify `300` .
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// The value against which the specified statistic is compared.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// A CloudWatch metric dimension.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The number of periods, in five-minute increments, during which the alarm condition must exist before the alarm triggers automatic scaling activity.
	//
	// The default value is `1` .
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The namespace for the CloudWatch metric.
	//
	// The default is `AWS/ElasticMapReduce` .
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The statistic to apply to the metric associated with the alarm.
	//
	// The default is `AVERAGE` .
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// The unit of measure associated with the CloudWatch metric being watched.
	//
	// The value specified for `Unit` must correspond to the units specified in the CloudWatch metric.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}


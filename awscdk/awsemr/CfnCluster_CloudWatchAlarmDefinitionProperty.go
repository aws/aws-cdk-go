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
//   cloudWatchAlarmDefinitionProperty := &CloudWatchAlarmDefinitionProperty{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	MetricName: jsii.String("metricName"),
//   	Period: jsii.Number(123),
//   	Threshold: jsii.Number(123),
//
//   	// the properties below are optional
//   	Dimensions: []interface{}{
//   		&MetricDimensionProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	EvaluationPeriods: jsii.Number(123),
//   	Namespace: jsii.String("namespace"),
//   	Statistic: jsii.String("statistic"),
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchalarmdefinition.html
//
type CfnCluster_CloudWatchAlarmDefinitionProperty struct {
	// Determines how the metric specified by `MetricName` is compared to the value specified by `Threshold` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchalarmdefinition.html#cfn-emr-cluster-cloudwatchalarmdefinition-comparisonoperator
	//
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The name of the CloudWatch metric that is watched to determine an alarm condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchalarmdefinition.html#cfn-emr-cluster-cloudwatchalarmdefinition-metricname
	//
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The period, in seconds, over which the statistic is applied.
	//
	// CloudWatch metrics for Amazon EMR are emitted every five minutes (300 seconds), so if you specify a CloudWatch metric, specify `300` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchalarmdefinition.html#cfn-emr-cluster-cloudwatchalarmdefinition-period
	//
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// The value against which the specified statistic is compared.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchalarmdefinition.html#cfn-emr-cluster-cloudwatchalarmdefinition-threshold
	//
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// A CloudWatch metric dimension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchalarmdefinition.html#cfn-emr-cluster-cloudwatchalarmdefinition-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The number of periods, in five-minute increments, during which the alarm condition must exist before the alarm triggers automatic scaling activity.
	//
	// The default value is `1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchalarmdefinition.html#cfn-emr-cluster-cloudwatchalarmdefinition-evaluationperiods
	//
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The namespace for the CloudWatch metric.
	//
	// The default is `AWS/ElasticMapReduce` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchalarmdefinition.html#cfn-emr-cluster-cloudwatchalarmdefinition-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The statistic to apply to the metric associated with the alarm.
	//
	// The default is `AVERAGE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchalarmdefinition.html#cfn-emr-cluster-cloudwatchalarmdefinition-statistic
	//
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// The unit of measure associated with the CloudWatch metric being watched.
	//
	// The value specified for `Unit` must correspond to the units specified in the CloudWatch metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchalarmdefinition.html#cfn-emr-cluster-cloudwatchalarmdefinition-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}


package awsemr


// `ScalingTrigger` is a subproperty of the `ScalingRule` property type.
//
// `ScalingTrigger` determines the conditions that trigger an automatic scaling activity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingTriggerProperty := &ScalingTriggerProperty{
//   	CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		MetricName: jsii.String("metricName"),
//   		Period: jsii.Number(123),
//   		Threshold: jsii.Number(123),
//
//   		// the properties below are optional
//   		Dimensions: []interface{}{
//   			&MetricDimensionProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		EvaluationPeriods: jsii.Number(123),
//   		Namespace: jsii.String("namespace"),
//   		Statistic: jsii.String("statistic"),
//   		Unit: jsii.String("unit"),
//   	},
//   }
//
type CfnInstanceGroupConfig_ScalingTriggerProperty struct {
	// The definition of a CloudWatch metric alarm.
	//
	// When the defined alarm conditions are met along with other trigger parameters, scaling activity begins.
	CloudWatchAlarmDefinition interface{} `field:"required" json:"cloudWatchAlarmDefinition" yaml:"cloudWatchAlarmDefinition"`
}


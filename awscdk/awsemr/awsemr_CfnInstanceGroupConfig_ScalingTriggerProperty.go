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
//   scalingTriggerProperty := &scalingTriggerProperty{
//   	cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   		comparisonOperator: jsii.String("comparisonOperator"),
//   		metricName: jsii.String("metricName"),
//   		period: jsii.Number(123),
//   		threshold: jsii.Number(123),
//
//   		// the properties below are optional
//   		dimensions: []interface{}{
//   			&metricDimensionProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		evaluationPeriods: jsii.Number(123),
//   		namespace: jsii.String("namespace"),
//   		statistic: jsii.String("statistic"),
//   		unit: jsii.String("unit"),
//   	},
//   }
//
type CfnInstanceGroupConfig_ScalingTriggerProperty struct {
	// The definition of a CloudWatch metric alarm.
	//
	// When the defined alarm conditions are met along with other trigger parameters, scaling activity begins.
	CloudWatchAlarmDefinition interface{} `field:"required" json:"cloudWatchAlarmDefinition" yaml:"cloudWatchAlarmDefinition"`
}


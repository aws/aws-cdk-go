package mixinsawsemr


// `ScalingTrigger` is a subproperty of the `ScalingRule` property type.
//
// `ScalingTrigger` determines the conditions that trigger an automatic scaling activity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scalingTriggerProperty := &ScalingTriggerProperty{
//   	CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		Dimensions: []interface{}{
//   			&MetricDimensionProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		EvaluationPeriods: jsii.Number(123),
//   		MetricName: jsii.String("metricName"),
//   		Namespace: jsii.String("namespace"),
//   		Period: jsii.Number(123),
//   		Statistic: jsii.String("statistic"),
//   		Threshold: jsii.Number(123),
//   		Unit: jsii.String("unit"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-scalingtrigger.html
//
type CfnClusterPropsMixin_ScalingTriggerProperty struct {
	// The definition of a CloudWatch metric alarm.
	//
	// When the defined alarm conditions are met along with other trigger parameters, scaling activity begins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-scalingtrigger.html#cfn-emr-cluster-scalingtrigger-cloudwatchalarmdefinition
	//
	CloudWatchAlarmDefinition interface{} `field:"optional" json:"cloudWatchAlarmDefinition" yaml:"cloudWatchAlarmDefinition"`
}


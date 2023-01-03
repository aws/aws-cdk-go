package awscloudwatch


// The `Metric` property type represents a specific metric.
//
// `Metric` is a property of the [MetricStat](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricstat.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricProperty := &metricProperty{
//   	dimensions: []interface{}{
//   		&dimensionProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	metricName: jsii.String("metricName"),
//   	namespace: jsii.String("namespace"),
//   }
//
type CfnAlarm_MetricProperty struct {
	// The metric dimensions that you want to be used for the metric that the alarm will watch..
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The name of the metric that you want the alarm to watch.
	//
	// This is a required field.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The namespace of the metric that the alarm will watch.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}


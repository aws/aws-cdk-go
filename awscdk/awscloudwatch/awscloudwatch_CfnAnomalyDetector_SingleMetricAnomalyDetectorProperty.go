package awscloudwatch


// Designates the CloudWatch metric and statistic that provides the time series the anomaly detector uses as input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   singleMetricAnomalyDetectorProperty := &singleMetricAnomalyDetectorProperty{
//   	dimensions: []interface{}{
//   		&dimensionProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	metricName: jsii.String("metricName"),
//   	namespace: jsii.String("namespace"),
//   	stat: jsii.String("stat"),
//   }
//
type CfnAnomalyDetector_SingleMetricAnomalyDetectorProperty struct {
	// The metric dimensions to create the anomaly detection model for.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The name of the metric to create the anomaly detection model for.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The namespace of the metric to create the anomaly detection model for.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The statistic to use for the metric and anomaly detection model.
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
}


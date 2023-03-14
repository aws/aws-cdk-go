package awscloudwatch


// Represents a specific metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricProperty := &metricProperty{
//   	metricName: jsii.String("metricName"),
//   	namespace: jsii.String("namespace"),
//
//   	// the properties below are optional
//   	dimensions: []interface{}{
//   		&dimensionProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAnomalyDetector_MetricProperty struct {
	// The name of the metric.
	//
	// This is a required field.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The namespace of the metric.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The dimensions for the metric.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
}


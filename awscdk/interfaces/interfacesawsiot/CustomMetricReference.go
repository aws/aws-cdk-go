package interfacesawsiot


// A reference to a CustomMetric resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customMetricReference := &CustomMetricReference{
//   	MetricName: jsii.String("metricName"),
//   }
//
type CustomMetricReference struct {
	// The MetricName of the CustomMetric resource.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
}


package interfacesawsconnect


// A reference to a Metric resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricReference := &MetricReference{
//   	MetricArn: jsii.String("metricArn"),
//   }
//
type MetricReference struct {
	// The MetricArn of the Metric resource.
	MetricArn *string `field:"required" json:"metricArn" yaml:"metricArn"`
}


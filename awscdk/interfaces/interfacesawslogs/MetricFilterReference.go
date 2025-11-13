package interfacesawslogs


// A reference to a MetricFilter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricFilterReference := &MetricFilterReference{
//   	FilterName: jsii.String("filterName"),
//   	LogGroupName: jsii.String("logGroupName"),
//   }
//
type MetricFilterReference struct {
	// The FilterName of the MetricFilter resource.
	FilterName *string `field:"required" json:"filterName" yaml:"filterName"`
	// The LogGroupName of the MetricFilter resource.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
}


package interfacesawsiot


// A reference to a FleetMetric resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fleetMetricReference := &FleetMetricReference{
//   	MetricName: jsii.String("metricName"),
//   }
//
type FleetMetricReference struct {
	// The MetricName of the FleetMetric resource.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
}


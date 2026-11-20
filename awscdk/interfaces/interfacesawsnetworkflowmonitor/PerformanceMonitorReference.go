package interfacesawsnetworkflowmonitor


// A reference to a PerformanceMonitor resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   performanceMonitorReference := &PerformanceMonitorReference{
//   	PerformanceMonitorArn: jsii.String("performanceMonitorArn"),
//   }
//
type PerformanceMonitorReference struct {
	// The Arn of the PerformanceMonitor resource.
	PerformanceMonitorArn *string `field:"required" json:"performanceMonitorArn" yaml:"performanceMonitorArn"`
}


package interfacesawsce


// A reference to a AnomalyMonitor resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   anomalyMonitorReference := &AnomalyMonitorReference{
//   	MonitorArn: jsii.String("monitorArn"),
//   }
//
type AnomalyMonitorReference struct {
	// The MonitorArn of the AnomalyMonitor resource.
	MonitorArn *string `field:"required" json:"monitorArn" yaml:"monitorArn"`
}


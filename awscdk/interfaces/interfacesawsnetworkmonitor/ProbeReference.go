package interfacesawsnetworkmonitor


// A reference to a Probe resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   probeReference := &ProbeReference{
//   	MonitorName: jsii.String("monitorName"),
//   	ProbeArn: jsii.String("probeArn"),
//   	ProbeId: jsii.String("probeId"),
//   }
//
type ProbeReference struct {
	// The MonitorName of the Probe resource.
	MonitorName *string `field:"required" json:"monitorName" yaml:"monitorName"`
	// The ARN of the Probe resource.
	ProbeArn *string `field:"required" json:"probeArn" yaml:"probeArn"`
	// The ProbeId of the Probe resource.
	ProbeId *string `field:"required" json:"probeId" yaml:"probeId"`
}


package awsinternetmonitor


// A reference to a Monitor resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitorReference := &MonitorReference{
//   	MonitorArn: jsii.String("monitorArn"),
//   	MonitorName: jsii.String("monitorName"),
//   }
//
type MonitorReference struct {
	// The ARN of the Monitor resource.
	MonitorArn *string `field:"required" json:"monitorArn" yaml:"monitorArn"`
	// The MonitorName of the Monitor resource.
	MonitorName *string `field:"required" json:"monitorName" yaml:"monitorName"`
}


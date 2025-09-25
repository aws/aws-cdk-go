package awssagemaker


// A reference to a MonitoringSchedule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringScheduleReference := &MonitoringScheduleReference{
//   	MonitoringScheduleArn: jsii.String("monitoringScheduleArn"),
//   }
//
type MonitoringScheduleReference struct {
	// The MonitoringScheduleArn of the MonitoringSchedule resource.
	MonitoringScheduleArn *string `field:"required" json:"monitoringScheduleArn" yaml:"monitoringScheduleArn"`
}


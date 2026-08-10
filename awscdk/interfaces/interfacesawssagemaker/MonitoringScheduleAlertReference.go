package interfacesawssagemaker


// A reference to a MonitoringScheduleAlert resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringScheduleAlertReference := &MonitoringScheduleAlertReference{
//   	MonitoringScheduleAlertArn: jsii.String("monitoringScheduleAlertArn"),
//   }
//
type MonitoringScheduleAlertReference struct {
	// The Arn of the MonitoringScheduleAlert resource.
	MonitoringScheduleAlertArn *string `field:"required" json:"monitoringScheduleAlertArn" yaml:"monitoringScheduleAlertArn"`
}


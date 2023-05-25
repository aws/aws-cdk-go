package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceUpdateSettingsProperty := &MaintenanceUpdateSettingsProperty{
//   	MaintenanceDay: jsii.String("maintenanceDay"),
//   	MaintenanceScheduledDate: jsii.String("maintenanceScheduledDate"),
//   	MaintenanceStartTime: jsii.String("maintenanceStartTime"),
//   }
//
type CfnChannel_MaintenanceUpdateSettingsProperty struct {
	// `CfnChannel.MaintenanceUpdateSettingsProperty.MaintenanceDay`.
	MaintenanceDay *string `field:"optional" json:"maintenanceDay" yaml:"maintenanceDay"`
	// `CfnChannel.MaintenanceUpdateSettingsProperty.MaintenanceScheduledDate`.
	MaintenanceScheduledDate *string `field:"optional" json:"maintenanceScheduledDate" yaml:"maintenanceScheduledDate"`
	// `CfnChannel.MaintenanceUpdateSettingsProperty.MaintenanceStartTime`.
	MaintenanceStartTime *string `field:"optional" json:"maintenanceStartTime" yaml:"maintenanceStartTime"`
}


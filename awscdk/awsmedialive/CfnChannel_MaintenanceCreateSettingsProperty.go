package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceCreateSettingsProperty := &MaintenanceCreateSettingsProperty{
//   	MaintenanceDay: jsii.String("maintenanceDay"),
//   	MaintenanceStartTime: jsii.String("maintenanceStartTime"),
//   }
//
type CfnChannel_MaintenanceCreateSettingsProperty struct {
	// `CfnChannel.MaintenanceCreateSettingsProperty.MaintenanceDay`.
	MaintenanceDay *string `field:"optional" json:"maintenanceDay" yaml:"maintenanceDay"`
	// `CfnChannel.MaintenanceCreateSettingsProperty.MaintenanceStartTime`.
	MaintenanceStartTime *string `field:"optional" json:"maintenanceStartTime" yaml:"maintenanceStartTime"`
}


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-maintenanceupdatesettings.html
//
type CfnChannel_MaintenanceUpdateSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-maintenanceupdatesettings.html#cfn-medialive-channel-maintenanceupdatesettings-maintenanceday
	//
	MaintenanceDay *string `field:"optional" json:"maintenanceDay" yaml:"maintenanceDay"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-maintenanceupdatesettings.html#cfn-medialive-channel-maintenanceupdatesettings-maintenancescheduleddate
	//
	MaintenanceScheduledDate *string `field:"optional" json:"maintenanceScheduledDate" yaml:"maintenanceScheduledDate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-maintenanceupdatesettings.html#cfn-medialive-channel-maintenanceupdatesettings-maintenancestarttime
	//
	MaintenanceStartTime *string `field:"optional" json:"maintenanceStartTime" yaml:"maintenanceStartTime"`
}


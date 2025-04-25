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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-maintenancecreatesettings.html
//
type CfnChannel_MaintenanceCreateSettingsProperty struct {
	// Choose one day of the week for maintenance.
	//
	// The chosen day is used for all future maintenance windows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-maintenancecreatesettings.html#cfn-medialive-channel-maintenancecreatesettings-maintenanceday
	//
	MaintenanceDay *string `field:"optional" json:"maintenanceDay" yaml:"maintenanceDay"`
	// Choose the hour that maintenance will start.
	//
	// The chosen time is used for all future maintenance windows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-maintenancecreatesettings.html#cfn-medialive-channel-maintenancecreatesettings-maintenancestarttime
	//
	MaintenanceStartTime *string `field:"optional" json:"maintenanceStartTime" yaml:"maintenanceStartTime"`
}


package awskinesisanalyticsv2


// Describes the maintenance configuration for the application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationMaintenanceConfigurationProperty := &ApplicationMaintenanceConfigurationProperty{
//   	ApplicationMaintenanceWindowStartTime: jsii.String("applicationMaintenanceWindowStartTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationmaintenanceconfiguration.html
//
type CfnApplication_ApplicationMaintenanceConfigurationProperty struct {
	// The start time for the maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationmaintenanceconfiguration.html#cfn-kinesisanalyticsv2-application-applicationmaintenanceconfiguration-applicationmaintenancewindowstarttime
	//
	ApplicationMaintenanceWindowStartTime *string `field:"required" json:"applicationMaintenanceWindowStartTime" yaml:"applicationMaintenanceWindowStartTime"`
}


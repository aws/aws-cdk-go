package awskinesisanalyticsv2


// Specifies the maintenance configuration for a AKAlong .
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
	// The UTC timestamp of a day from which the eight-hour maintenance window will begin every day of the week.
	//
	// Maintenance of the application happens only during this eight-hour window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationmaintenanceconfiguration.html#cfn-kinesisanalyticsv2-application-applicationmaintenanceconfiguration-applicationmaintenancewindowstarttime
	//
	ApplicationMaintenanceWindowStartTime *string `field:"required" json:"applicationMaintenanceWindowStartTime" yaml:"applicationMaintenanceWindowStartTime"`
}


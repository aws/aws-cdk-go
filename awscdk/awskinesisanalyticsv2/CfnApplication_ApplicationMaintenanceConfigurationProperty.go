package awskinesisanalyticsv2


// Specifies the maintence window parameters for a Kinesis Data Analytics application.
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
type CfnApplication_ApplicationMaintenanceConfigurationProperty struct {
	// Specifies the start time of the maintence window.
	ApplicationMaintenanceWindowStartTime *string `field:"required" json:"applicationMaintenanceWindowStartTime" yaml:"applicationMaintenanceWindowStartTime"`
}


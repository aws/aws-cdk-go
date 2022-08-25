package awskinesisanalyticsv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationMaintenanceConfigurationProperty := &applicationMaintenanceConfigurationProperty{
//   	applicationMaintenanceWindowStartTime: jsii.String("applicationMaintenanceWindowStartTime"),
//   }
//
type CfnApplication_ApplicationMaintenanceConfigurationProperty struct {
	// `CfnApplication.ApplicationMaintenanceConfigurationProperty.ApplicationMaintenanceWindowStartTime`.
	ApplicationMaintenanceWindowStartTime *string `field:"required" json:"applicationMaintenanceWindowStartTime" yaml:"applicationMaintenanceWindowStartTime"`
}


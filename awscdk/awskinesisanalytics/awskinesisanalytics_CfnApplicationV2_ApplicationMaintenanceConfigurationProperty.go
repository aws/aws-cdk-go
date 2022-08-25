package awskinesisanalytics


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationMaintenanceConfigurationProperty := &applicationMaintenanceConfigurationProperty{
//   	applicationMaintenanceWindowStartTime: jsii.String("applicationMaintenanceWindowStartTime"),
//   }
//
type CfnApplicationV2_ApplicationMaintenanceConfigurationProperty struct {
	// `CfnApplicationV2.ApplicationMaintenanceConfigurationProperty.ApplicationMaintenanceWindowStartTime`.
	ApplicationMaintenanceWindowStartTime *string `field:"required" json:"applicationMaintenanceWindowStartTime" yaml:"applicationMaintenanceWindowStartTime"`
}


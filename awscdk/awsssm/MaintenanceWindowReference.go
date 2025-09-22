package awsssm


// A reference to a MaintenanceWindow resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceWindowReference := &MaintenanceWindowReference{
//   	MaintenanceWindowId: jsii.String("maintenanceWindowId"),
//   }
//
type MaintenanceWindowReference struct {
	// The Id of the MaintenanceWindow resource.
	MaintenanceWindowId *string `field:"required" json:"maintenanceWindowId" yaml:"maintenanceWindowId"`
}


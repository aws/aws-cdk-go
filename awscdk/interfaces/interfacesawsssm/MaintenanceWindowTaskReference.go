package interfacesawsssm


// A reference to a MaintenanceWindowTask resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceWindowTaskReference := &MaintenanceWindowTaskReference{
//   	MaintenanceWindowTaskId: jsii.String("maintenanceWindowTaskId"),
//   	WindowTaskId: jsii.String("windowTaskId"),
//   }
//
type MaintenanceWindowTaskReference struct {
	// The Id of the MaintenanceWindowTask resource.
	MaintenanceWindowTaskId *string `field:"required" json:"maintenanceWindowTaskId" yaml:"maintenanceWindowTaskId"`
	// The WindowTaskId of the MaintenanceWindowTask resource.
	WindowTaskId *string `field:"required" json:"windowTaskId" yaml:"windowTaskId"`
}


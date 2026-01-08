package interfacesawsssm


// A reference to a MaintenanceWindowTask resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceWindowTaskReference := &MaintenanceWindowTaskReference{
//   	WindowId: jsii.String("windowId"),
//   	WindowTaskId: jsii.String("windowTaskId"),
//   }
//
type MaintenanceWindowTaskReference struct {
	// The WindowId of the MaintenanceWindowTask resource.
	WindowId *string `field:"required" json:"windowId" yaml:"windowId"`
	// The WindowTaskId of the MaintenanceWindowTask resource.
	WindowTaskId *string `field:"required" json:"windowTaskId" yaml:"windowTaskId"`
}


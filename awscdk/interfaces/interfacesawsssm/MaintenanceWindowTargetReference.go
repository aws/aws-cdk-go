package interfacesawsssm


// A reference to a MaintenanceWindowTarget resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceWindowTargetReference := &MaintenanceWindowTargetReference{
//   	WindowId: jsii.String("windowId"),
//   	WindowTargetId: jsii.String("windowTargetId"),
//   }
//
type MaintenanceWindowTargetReference struct {
	// The WindowId of the MaintenanceWindowTarget resource.
	WindowId *string `field:"required" json:"windowId" yaml:"windowId"`
	// The WindowTargetId of the MaintenanceWindowTarget resource.
	WindowTargetId *string `field:"required" json:"windowTargetId" yaml:"windowTargetId"`
}


package interfacesawsssm


// A reference to a MaintenanceWindow resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceWindowReference := &MaintenanceWindowReference{
//   	WindowId: jsii.String("windowId"),
//   }
//
type MaintenanceWindowReference struct {
	// The WindowId of the MaintenanceWindow resource.
	WindowId *string `field:"required" json:"windowId" yaml:"windowId"`
}


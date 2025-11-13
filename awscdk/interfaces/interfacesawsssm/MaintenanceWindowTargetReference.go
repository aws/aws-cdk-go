package interfacesawsssm


// A reference to a MaintenanceWindowTarget resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceWindowTargetReference := &MaintenanceWindowTargetReference{
//   	MaintenanceWindowTargetId: jsii.String("maintenanceWindowTargetId"),
//   }
//
type MaintenanceWindowTargetReference struct {
	// The Id of the MaintenanceWindowTarget resource.
	MaintenanceWindowTargetId *string `field:"required" json:"maintenanceWindowTargetId" yaml:"maintenanceWindowTargetId"`
}


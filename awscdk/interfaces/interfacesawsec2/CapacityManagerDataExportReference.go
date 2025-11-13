package interfacesawsec2


// A reference to a CapacityManagerDataExport resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityManagerDataExportReference := &CapacityManagerDataExportReference{
//   	CapacityManagerDataExportId: jsii.String("capacityManagerDataExportId"),
//   }
//
type CapacityManagerDataExportReference struct {
	// The CapacityManagerDataExportId of the CapacityManagerDataExport resource.
	CapacityManagerDataExportId *string `field:"required" json:"capacityManagerDataExportId" yaml:"capacityManagerDataExportId"`
}


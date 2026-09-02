package interfacesawsinvoicing


// A reference to a ProcurementPortalPreference resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   procurementPortalPreferenceReference := &ProcurementPortalPreferenceReference{
//   	ProcurementPortalPreferenceArn: jsii.String("procurementPortalPreferenceArn"),
//   }
//
type ProcurementPortalPreferenceReference struct {
	// The ProcurementPortalPreferenceArn of the ProcurementPortalPreference resource.
	ProcurementPortalPreferenceArn *string `field:"required" json:"procurementPortalPreferenceArn" yaml:"procurementPortalPreferenceArn"`
}


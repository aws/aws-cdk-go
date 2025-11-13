package interfacesawsses


// A reference to a VdmAttributes resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vdmAttributesReference := &VdmAttributesReference{
//   	VdmAttributesResourceId: jsii.String("vdmAttributesResourceId"),
//   }
//
type VdmAttributesReference struct {
	// The VdmAttributesResourceId of the VdmAttributes resource.
	VdmAttributesResourceId *string `field:"required" json:"vdmAttributesResourceId" yaml:"vdmAttributesResourceId"`
}


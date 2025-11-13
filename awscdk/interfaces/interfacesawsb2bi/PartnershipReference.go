package interfacesawsb2bi


// A reference to a Partnership resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partnershipReference := &PartnershipReference{
//   	PartnershipArn: jsii.String("partnershipArn"),
//   	PartnershipId: jsii.String("partnershipId"),
//   }
//
type PartnershipReference struct {
	// The ARN of the Partnership resource.
	PartnershipArn *string `field:"required" json:"partnershipArn" yaml:"partnershipArn"`
	// The PartnershipId of the Partnership resource.
	PartnershipId *string `field:"required" json:"partnershipId" yaml:"partnershipId"`
}


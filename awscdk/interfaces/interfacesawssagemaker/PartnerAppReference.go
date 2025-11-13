package interfacesawssagemaker


// A reference to a PartnerApp resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partnerAppReference := &PartnerAppReference{
//   	PartnerAppArn: jsii.String("partnerAppArn"),
//   }
//
type PartnerAppReference struct {
	// The Arn of the PartnerApp resource.
	PartnerAppArn *string `field:"required" json:"partnerAppArn" yaml:"partnerAppArn"`
}


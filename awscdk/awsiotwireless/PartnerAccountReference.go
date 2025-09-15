package awsiotwireless


// A reference to a PartnerAccount resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partnerAccountReference := &PartnerAccountReference{
//   	PartnerAccountArn: jsii.String("partnerAccountArn"),
//   	PartnerAccountId: jsii.String("partnerAccountId"),
//   }
//
type PartnerAccountReference struct {
	// The ARN of the PartnerAccount resource.
	PartnerAccountArn *string `field:"required" json:"partnerAccountArn" yaml:"partnerAccountArn"`
	// The PartnerAccountId of the PartnerAccount resource.
	PartnerAccountId *string `field:"required" json:"partnerAccountId" yaml:"partnerAccountId"`
}


package awselasticloadbalancingv2


// A reference to a TrustStoreRevocation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trustStoreRevocationReference := &TrustStoreRevocationReference{
//   	RevocationId: jsii.String("revocationId"),
//   	TrustStoreArn: jsii.String("trustStoreArn"),
//   }
//
type TrustStoreRevocationReference struct {
	// The RevocationId of the TrustStoreRevocation resource.
	RevocationId *string `field:"required" json:"revocationId" yaml:"revocationId"`
	// The TrustStoreArn of the TrustStoreRevocation resource.
	TrustStoreArn *string `field:"required" json:"trustStoreArn" yaml:"trustStoreArn"`
}


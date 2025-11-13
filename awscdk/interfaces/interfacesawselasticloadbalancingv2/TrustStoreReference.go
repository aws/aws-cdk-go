package interfacesawselasticloadbalancingv2


// A reference to a TrustStore resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trustStoreReference := &TrustStoreReference{
//   	TrustStoreArn: jsii.String("trustStoreArn"),
//   }
//
type TrustStoreReference struct {
	// The TrustStoreArn of the TrustStore resource.
	TrustStoreArn *string `field:"required" json:"trustStoreArn" yaml:"trustStoreArn"`
}


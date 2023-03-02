package awslicensemanager


// Details associated with the issuer of a license.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   issuerDataProperty := &issuerDataProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	signKey: jsii.String("signKey"),
//   }
//
type CfnLicense_IssuerDataProperty struct {
	// Issuer name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Asymmetric KMS key from AWS Key Management Service .
	//
	// The KMS key must have a key usage of sign and verify, and support the RSASSA-PSS SHA-256 signing algorithm.
	SignKey *string `field:"optional" json:"signKey" yaml:"signKey"`
}


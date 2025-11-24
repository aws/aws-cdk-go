package mixinsawslicensemanager


// Details associated with the issuer of a license.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   issuerDataProperty := &IssuerDataProperty{
//   	Name: jsii.String("name"),
//   	SignKey: jsii.String("signKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-issuerdata.html
//
type CfnLicensePropsMixin_IssuerDataProperty struct {
	// Issuer name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-issuerdata.html#cfn-licensemanager-license-issuerdata-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Asymmetric KMS key from AWS Key Management Service .
	//
	// The KMS key must have a key usage of sign and verify, and support the RSASSA-PSS SHA-256 signing algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-issuerdata.html#cfn-licensemanager-license-issuerdata-signkey
	//
	SignKey *string `field:"optional" json:"signKey" yaml:"signKey"`
}


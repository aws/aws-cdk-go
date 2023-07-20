package awsacmpca


// Specifies additional purposes for which the certified public key may be used other than basic purposes indicated in the `KeyUsage` extension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   extendedKeyUsageProperty := &ExtendedKeyUsageProperty{
//   	ExtendedKeyUsageObjectIdentifier: jsii.String("extendedKeyUsageObjectIdentifier"),
//   	ExtendedKeyUsageType: jsii.String("extendedKeyUsageType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extendedkeyusage.html
//
type CfnCertificate_ExtendedKeyUsageProperty struct {
	// Specifies a custom `ExtendedKeyUsage` with an object identifier (OID).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extendedkeyusage.html#cfn-acmpca-certificate-extendedkeyusage-extendedkeyusageobjectidentifier
	//
	ExtendedKeyUsageObjectIdentifier *string `field:"optional" json:"extendedKeyUsageObjectIdentifier" yaml:"extendedKeyUsageObjectIdentifier"`
	// Specifies a standard `ExtendedKeyUsage` as defined as in [RFC 5280](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extendedkeyusage.html#cfn-acmpca-certificate-extendedkeyusage-extendedkeyusagetype
	//
	ExtendedKeyUsageType *string `field:"optional" json:"extendedKeyUsageType" yaml:"extendedKeyUsageType"`
}


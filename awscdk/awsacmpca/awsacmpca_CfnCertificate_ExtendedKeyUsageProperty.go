package awsacmpca


// Specifies additional purposes for which the certified public key may be used other than basic purposes indicated in the `KeyUsage` extension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   extendedKeyUsageProperty := &extendedKeyUsageProperty{
//   	extendedKeyUsageObjectIdentifier: jsii.String("extendedKeyUsageObjectIdentifier"),
//   	extendedKeyUsageType: jsii.String("extendedKeyUsageType"),
//   }
//
type CfnCertificate_ExtendedKeyUsageProperty struct {
	// Specifies a custom `ExtendedKeyUsage` with an object identifier (OID).
	ExtendedKeyUsageObjectIdentifier *string `field:"optional" json:"extendedKeyUsageObjectIdentifier" yaml:"extendedKeyUsageObjectIdentifier"`
	// Specifies a standard `ExtendedKeyUsage` as defined as in [RFC 5280](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12) .
	ExtendedKeyUsageType *string `field:"optional" json:"extendedKeyUsageType" yaml:"extendedKeyUsageType"`
}


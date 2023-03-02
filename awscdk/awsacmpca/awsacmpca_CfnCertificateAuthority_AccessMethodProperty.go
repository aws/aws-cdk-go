package awsacmpca


// Describes the type and format of extension access.
//
// Only one of `CustomObjectIdentifier` or `AccessMethodType` may be provided. Providing both results in `InvalidArgsException` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessMethodProperty := &accessMethodProperty{
//   	accessMethodType: jsii.String("accessMethodType"),
//   	customObjectIdentifier: jsii.String("customObjectIdentifier"),
//   }
//
type CfnCertificateAuthority_AccessMethodProperty struct {
	// Specifies the `AccessMethod` .
	AccessMethodType *string `field:"optional" json:"accessMethodType" yaml:"accessMethodType"`
	// An object identifier (OID) specifying the `AccessMethod` .
	//
	// The OID must satisfy the regular expression shown below. For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
	CustomObjectIdentifier *string `field:"optional" json:"customObjectIdentifier" yaml:"customObjectIdentifier"`
}


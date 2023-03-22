package awsacmpca


// Defines a custom ASN.1 X.400 `GeneralName` using an object identifier (OID) and value. The OID must satisfy the regular expression shown below. For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   otherNameProperty := &otherNameProperty{
//   	typeId: jsii.String("typeId"),
//   	value: jsii.String("value"),
//   }
//
type CfnCertificate_OtherNameProperty struct {
	// Specifies an OID.
	TypeId *string `field:"required" json:"typeId" yaml:"typeId"`
	// Specifies an OID value.
	Value *string `field:"required" json:"value" yaml:"value"`
}


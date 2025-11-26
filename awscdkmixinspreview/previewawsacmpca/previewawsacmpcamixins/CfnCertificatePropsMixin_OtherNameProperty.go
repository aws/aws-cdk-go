package previewawsacmpcamixins


// Defines a custom ASN.1 X.400 `GeneralName` using an object identifier (OID) and value. The OID must satisfy the regular expression shown below. For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   otherNameProperty := &OtherNameProperty{
//   	TypeId: jsii.String("typeId"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-othername.html
//
type CfnCertificatePropsMixin_OtherNameProperty struct {
	// Specifies an OID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-othername.html#cfn-acmpca-certificate-othername-typeid
	//
	TypeId *string `field:"optional" json:"typeId" yaml:"typeId"`
	// Specifies an OID value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-othername.html#cfn-acmpca-certificate-othername-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}


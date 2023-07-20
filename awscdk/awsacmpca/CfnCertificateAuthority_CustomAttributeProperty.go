package awsacmpca


// Defines the X.500 relative distinguished name (RDN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customAttributeProperty := &CustomAttributeProperty{
//   	ObjectIdentifier: jsii.String("objectIdentifier"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-customattribute.html
//
type CfnCertificateAuthority_CustomAttributeProperty struct {
	// Specifies the object identifier (OID) of the attribute type of the relative distinguished name (RDN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-customattribute.html#cfn-acmpca-certificateauthority-customattribute-objectidentifier
	//
	ObjectIdentifier *string `field:"required" json:"objectIdentifier" yaml:"objectIdentifier"`
	// Specifies the attribute value of relative distinguished name (RDN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-customattribute.html#cfn-acmpca-certificateauthority-customattribute-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}


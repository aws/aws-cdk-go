package awsacmpca


// Defines the X.500 relative distinguished name (RDN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customAttributeProperty := &customAttributeProperty{
//   	objectIdentifier: jsii.String("objectIdentifier"),
//   	value: jsii.String("value"),
//   }
//
type CfnCertificate_CustomAttributeProperty struct {
	// Specifies the object identifier (OID) of the attribute type of the relative distinguished name (RDN).
	ObjectIdentifier *string `field:"required" json:"objectIdentifier" yaml:"objectIdentifier"`
	// Specifies the attribute value of relative distinguished name (RDN).
	Value *string `field:"required" json:"value" yaml:"value"`
}


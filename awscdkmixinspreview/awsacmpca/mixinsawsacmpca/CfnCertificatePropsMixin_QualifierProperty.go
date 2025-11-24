package mixinsawsacmpca


// Defines a `PolicyInformation` qualifier.
//
// AWS Private CA supports the [certification practice statement (CPS) qualifier](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.4) defined in RFC 5280.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   qualifierProperty := &QualifierProperty{
//   	CpsUri: jsii.String("cpsUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-qualifier.html
//
type CfnCertificatePropsMixin_QualifierProperty struct {
	// Contains a pointer to a certification practice statement (CPS) published by the CA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-qualifier.html#cfn-acmpca-certificate-qualifier-cpsuri
	//
	CpsUri *string `field:"optional" json:"cpsUri" yaml:"cpsUri"`
}


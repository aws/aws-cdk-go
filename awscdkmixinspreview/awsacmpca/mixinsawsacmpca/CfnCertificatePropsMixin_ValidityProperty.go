package mixinsawsacmpca


// Length of time for which the certificate issued by your private certificate authority (CA), or by the private CA itself, is valid in days, months, or years.
//
// You can issue a certificate by calling the `IssueCertificate` operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   validityProperty := &ValidityProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-validity.html
//
type CfnCertificatePropsMixin_ValidityProperty struct {
	// Specifies whether the `Value` parameter represents days, months, or years.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-validity.html#cfn-acmpca-certificate-validity-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// A long integer interpreted according to the value of `Type` , below.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-validity.html#cfn-acmpca-certificate-validity-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}


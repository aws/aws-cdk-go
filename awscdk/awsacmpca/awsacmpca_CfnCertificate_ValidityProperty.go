package awsacmpca


// Length of time for which the certificate issued by your private certificate authority (CA), or by the private CA itself, is valid in days, months, or years.
//
// You can issue a certificate by calling the `IssueCertificate` operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   validityProperty := &validityProperty{
//   	type: jsii.String("type"),
//   	value: jsii.Number(123),
//   }
//
type CfnCertificate_ValidityProperty struct {
	// Specifies whether the `Value` parameter represents days, months, or years.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A long integer interpreted according to the value of `Type` , below.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}


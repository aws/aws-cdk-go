package awssigner


// The validity period for the signing job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signatureValidityPeriodProperty := &signatureValidityPeriodProperty{
//   	type: jsii.String("type"),
//   	value: jsii.Number(123),
//   }
//
type CfnSigningProfile_SignatureValidityPeriodProperty struct {
	// The time unit for signature validity: DAYS | MONTHS | YEARS.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The numerical value of the time unit for signature validity.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}


package awssigner


// The validity period for the signing job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signatureValidityPeriodProperty := &SignatureValidityPeriodProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-signer-signingprofile-signaturevalidityperiod.html
//
type CfnSigningProfile_SignatureValidityPeriodProperty struct {
	// The time unit for signature validity: DAYS | MONTHS | YEARS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-signer-signingprofile-signaturevalidityperiod.html#cfn-signer-signingprofile-signaturevalidityperiod-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The numerical value of the time unit for signature validity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-signer-signingprofile-signaturevalidityperiod.html#cfn-signer-signingprofile-signaturevalidityperiod-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}


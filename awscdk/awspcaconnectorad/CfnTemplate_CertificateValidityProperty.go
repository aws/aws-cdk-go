package awspcaconnectorad


// Information describing the end of the validity period of the certificate.
//
// This parameter sets the “Not After” date for the certificate. Certificate validity is the period of time during which a certificate is valid. Validity can be expressed as an explicit date and time when the certificate expires, or as a span of time after issuance, stated in days, months, or years. For more information, see Validity in RFC 5280. This value is unaffected when ValidityNotBefore is also specified. For example, if Validity is set to 20 days in the future, the certificate will expire 20 days from issuance time regardless of the ValidityNotBefore value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateValidityProperty := &CertificateValidityProperty{
//   	RenewalPeriod: &ValidityPeriodProperty{
//   		Period: jsii.Number(123),
//   		PeriodType: jsii.String("periodType"),
//   	},
//   	ValidityPeriod: &ValidityPeriodProperty{
//   		Period: jsii.Number(123),
//   		PeriodType: jsii.String("periodType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-certificatevalidity.html
//
type CfnTemplate_CertificateValidityProperty struct {
	// Renewal period is the period of time before certificate expiration when a new certificate will be requested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-certificatevalidity.html#cfn-pcaconnectorad-template-certificatevalidity-renewalperiod
	//
	RenewalPeriod interface{} `field:"required" json:"renewalPeriod" yaml:"renewalPeriod"`
	// Information describing the end of the validity period of the certificate.
	//
	// This parameter sets the “Not After” date for the certificate. Certificate validity is the period of time during which a certificate is valid. Validity can be expressed as an explicit date and time when the certificate expires, or as a span of time after issuance, stated in days, months, or years. For more information, see Validity in RFC 5280. This value is unaffected when ValidityNotBefore is also specified. For example, if Validity is set to 20 days in the future, the certificate will expire 20 days from issuance time regardless of the ValidityNotBefore value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-certificatevalidity.html#cfn-pcaconnectorad-template-certificatevalidity-validityperiod
	//
	ValidityPeriod interface{} `field:"required" json:"validityPeriod" yaml:"validityPeriod"`
}


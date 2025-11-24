package mixinsawspcaconnectorad


// Information describing the end of the validity period of the certificate.
//
// This parameter sets the “Not After” date for the certificate. Certificate validity is the period of time during which a certificate is valid. Validity can be expressed as an explicit date and time when the certificate expires, or as a span of time after issuance, stated in hours, days, months, or years. For more information, see Validity in RFC 5280. This value is unaffected when ValidityNotBefore is also specified. For example, if Validity is set to 20 days in the future, the certificate will expire 20 days from issuance time regardless of the ValidityNotBefore value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   validityPeriodProperty := &ValidityPeriodProperty{
//   	Period: jsii.Number(123),
//   	PeriodType: jsii.String("periodType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-validityperiod.html
//
type CfnTemplatePropsMixin_ValidityPeriodProperty struct {
	// The numeric value for the validity period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-validityperiod.html#cfn-pcaconnectorad-template-validityperiod-period
	//
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The unit of time.
	//
	// You can select hours, days, weeks, months, and years.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-validityperiod.html#cfn-pcaconnectorad-template-validityperiod-periodtype
	//
	PeriodType *string `field:"optional" json:"periodType" yaml:"periodType"`
}


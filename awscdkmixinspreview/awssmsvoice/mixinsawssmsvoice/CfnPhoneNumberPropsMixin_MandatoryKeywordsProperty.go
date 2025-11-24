package mixinsawssmsvoice


// The keywords `HELP` and `STOP` are mandatory keywords that each phone number must have.
//
// For more information, see [Keywords](https://docs.aws.amazon.com/sms-voice/latest/userguide/keywords.html) in the End User Messaging  User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mandatoryKeywordsProperty := &MandatoryKeywordsProperty{
//   	Help: &MandatoryKeywordProperty{
//   		Message: jsii.String("message"),
//   	},
//   	Stop: &MandatoryKeywordProperty{
//   		Message: jsii.String("message"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-mandatorykeywords.html
//
type CfnPhoneNumberPropsMixin_MandatoryKeywordsProperty struct {
	// Specifies the `HELP` keyword that customers use to obtain customer support for this phone number.
	//
	// For more information, see [Keywords](https://docs.aws.amazon.com/sms-voice/latest/userguide/keywords.html) in the End User Messaging  User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-mandatorykeywords.html#cfn-smsvoice-phonenumber-mandatorykeywords-help
	//
	Help interface{} `field:"optional" json:"help" yaml:"help"`
	// Specifies the `STOP` keyword that customers use to opt out of receiving messages from this phone number.
	//
	// For more information, see [Required opt-out keywords](https://docs.aws.amazon.com/sms-voice/latest/userguide/keywords-required.html) in the End User Messaging  User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-mandatorykeywords.html#cfn-smsvoice-phonenumber-mandatorykeywords-stop
	//
	Stop interface{} `field:"optional" json:"stop" yaml:"stop"`
}


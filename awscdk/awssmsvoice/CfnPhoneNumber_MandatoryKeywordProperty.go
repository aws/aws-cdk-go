package awssmsvoice


// The keywords `HELP` and `STOP` are mandatory keywords that each phone number must have.
//
// For more information, see [Keywords](https://docs.aws.amazon.com/sms-voice/latest/userguide/keywords.html) in the End User Messaging  User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mandatoryKeywordProperty := &MandatoryKeywordProperty{
//   	Message: jsii.String("message"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-mandatorykeyword.html
//
type CfnPhoneNumber_MandatoryKeywordProperty struct {
	// The message associated with the keyword.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-mandatorykeyword.html#cfn-smsvoice-phonenumber-mandatorykeyword-message
	//
	Message *string `field:"required" json:"message" yaml:"message"`
}


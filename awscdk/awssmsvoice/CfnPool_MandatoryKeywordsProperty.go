package awssmsvoice


// The manadatory keywords, `HELP` and `STOP` to add to the pool.
//
// For more information, see [Keywords](https://docs.aws.amazon.com/sms-voice/latest/userguide/keywords.html) in the End User Messaging  User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-pool-mandatorykeywords.html
//
type CfnPool_MandatoryKeywordsProperty struct {
	// Specifies the pool's `HELP` keyword.
	//
	// For more information, see [Opt out list required keywords](https://docs.aws.amazon.com/sms-voice/latest/userguide/opt-out-list-keywords.html) in the End User Messaging  User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-pool-mandatorykeywords.html#cfn-smsvoice-pool-mandatorykeywords-help
	//
	Help interface{} `field:"required" json:"help" yaml:"help"`
	// Specifies the pool's opt-out keyword.
	//
	// For more information, see [Required opt-out keywords](https://docs.aws.amazon.com/sms-voice/latest/userguide/keywords-required.html) in the End User Messaging  User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-pool-mandatorykeywords.html#cfn-smsvoice-pool-mandatorykeywords-stop
	//
	Stop interface{} `field:"required" json:"stop" yaml:"stop"`
}


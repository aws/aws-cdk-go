package mixinsawslex


// Specifies the audio input specifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   audioSpecificationProperty := &AudioSpecificationProperty{
//   	EndTimeoutMs: jsii.Number(123),
//   	MaxLengthMs: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiospecification.html
//
type CfnBotPropsMixin_AudioSpecificationProperty struct {
	// Time for which a bot waits after the customer stops speaking to assume the utterance is finished.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiospecification.html#cfn-lex-bot-audiospecification-endtimeoutms
	//
	EndTimeoutMs *float64 `field:"optional" json:"endTimeoutMs" yaml:"endTimeoutMs"`
	// Time for how long Amazon Lex waits before speech input is truncated and the speech is returned to application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiospecification.html#cfn-lex-bot-audiospecification-maxlengthms
	//
	MaxLengthMs *float64 `field:"optional" json:"maxLengthMs" yaml:"maxLengthMs"`
}


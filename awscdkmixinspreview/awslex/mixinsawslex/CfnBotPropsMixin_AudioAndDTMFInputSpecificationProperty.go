package mixinsawslex


// Specifies the audio and DTMF input specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   audioAndDTMFInputSpecificationProperty := &AudioAndDTMFInputSpecificationProperty{
//   	AudioSpecification: &AudioSpecificationProperty{
//   		EndTimeoutMs: jsii.Number(123),
//   		MaxLengthMs: jsii.Number(123),
//   	},
//   	DtmfSpecification: &DTMFSpecificationProperty{
//   		DeletionCharacter: jsii.String("deletionCharacter"),
//   		EndCharacter: jsii.String("endCharacter"),
//   		EndTimeoutMs: jsii.Number(123),
//   		MaxLength: jsii.Number(123),
//   	},
//   	StartTimeoutMs: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audioanddtmfinputspecification.html
//
type CfnBotPropsMixin_AudioAndDTMFInputSpecificationProperty struct {
	// Specifies the settings on audio input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audioanddtmfinputspecification.html#cfn-lex-bot-audioanddtmfinputspecification-audiospecification
	//
	AudioSpecification interface{} `field:"optional" json:"audioSpecification" yaml:"audioSpecification"`
	// Specifies the settings on DTMF input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audioanddtmfinputspecification.html#cfn-lex-bot-audioanddtmfinputspecification-dtmfspecification
	//
	DtmfSpecification interface{} `field:"optional" json:"dtmfSpecification" yaml:"dtmfSpecification"`
	// Time for which a bot waits before assuming that the customer isn't going to speak or press a key.
	//
	// This timeout is shared between Audio and DTMF inputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audioanddtmfinputspecification.html#cfn-lex-bot-audioanddtmfinputspecification-starttimeoutms
	//
	StartTimeoutMs *float64 `field:"optional" json:"startTimeoutMs" yaml:"startTimeoutMs"`
}


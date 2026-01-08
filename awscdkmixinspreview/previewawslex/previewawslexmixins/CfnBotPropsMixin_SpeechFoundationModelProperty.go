package previewawslexmixins


// Configuration for a foundation model used for speech synthesis and recognition capabilities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   speechFoundationModelProperty := &SpeechFoundationModelProperty{
//   	ModelArn: jsii.String("modelArn"),
//   	VoiceId: jsii.String("voiceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-speechfoundationmodel.html
//
type CfnBotPropsMixin_SpeechFoundationModelProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-speechfoundationmodel.html#cfn-lex-bot-speechfoundationmodel-modelarn
	//
	ModelArn *string `field:"optional" json:"modelArn" yaml:"modelArn"`
	// The identifier of the voice to use for speech synthesis with the foundation model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-speechfoundationmodel.html#cfn-lex-bot-speechfoundationmodel-voiceid
	//
	VoiceId *string `field:"optional" json:"voiceId" yaml:"voiceId"`
}


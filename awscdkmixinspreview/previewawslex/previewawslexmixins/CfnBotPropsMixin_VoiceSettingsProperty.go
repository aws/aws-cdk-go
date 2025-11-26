package previewawslexmixins


// Defines settings for using an Amazon Polly voice to communicate with a user.
//
// Valid values include:
//
// - `standard`
// - `neural`
// - `long-form`
// - `generative`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceSettingsProperty := &VoiceSettingsProperty{
//   	Engine: jsii.String("engine"),
//   	VoiceId: jsii.String("voiceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-voicesettings.html
//
type CfnBotPropsMixin_VoiceSettingsProperty struct {
	// Indicates the type of Amazon Polly voice that Amazon Lex should use for voice interaction with the user.
	//
	// For more information, see the [`engine` parameter of the `SynthesizeSpeech` operation](https://docs.aws.amazon.com/polly/latest/dg/API_SynthesizeSpeech.html#polly-SynthesizeSpeech-request-Engine) in the *Amazon Polly developer guide* .
	//
	// If you do not specify a value, the default is `standard` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-voicesettings.html#cfn-lex-bot-voicesettings-engine
	//
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The identifier of the Amazon Polly voice to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-voicesettings.html#cfn-lex-bot-voicesettings-voiceid
	//
	VoiceId *string `field:"optional" json:"voiceId" yaml:"voiceId"`
}


package awslex


// Configures the Assisted Natural Language Understanding (NLU) feature for your bot.
//
// This specification determines whether enhanced intent recognition and utterance understanding capabilities are active.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nluImprovementSpecificationProperty := &NluImprovementSpecificationProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	AssistedNluMode: jsii.String("assistedNluMode"),
//   	IntentDisambiguationSettings: &IntentDisambiguationSettingsProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		CustomDisambiguationMessage: jsii.String("customDisambiguationMessage"),
//   		MaxDisambiguationIntents: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-nluimprovementspecification.html
//
type CfnBot_NluImprovementSpecificationProperty struct {
	// Determines whether the Assisted NLU feature is enabled for the bot.
	//
	// When set to `true` , Amazon Lex uses advanced models to improve intent recognition and slot resolution, with the default being `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-nluimprovementspecification.html#cfn-lex-bot-nluimprovementspecification-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the mode for Assisted NLU operation.
	//
	// Use `Primary` to make Assisted NLU the primary intent recognition method, or `Fallback` to use it only when standard NLU confidence is low.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-nluimprovementspecification.html#cfn-lex-bot-nluimprovementspecification-assistednlumode
	//
	AssistedNluMode *string `field:"optional" json:"assistedNluMode" yaml:"assistedNluMode"`
	// An object containing specifications for the Intent Disambiguation feature within the Assisted NLU settings.
	//
	// These settings determine how the bot handles ambiguous user inputs that could match multiple intents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-nluimprovementspecification.html#cfn-lex-bot-nluimprovementspecification-intentdisambiguationsettings
	//
	IntentDisambiguationSettings interface{} `field:"optional" json:"intentDisambiguationSettings" yaml:"intentDisambiguationSettings"`
}


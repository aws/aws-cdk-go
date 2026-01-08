package awslex


// Configures the Intent Disambiguation feature that helps resolve ambiguous user inputs when multiple intents could match.
//
// When enabled, the system presents clarifying questions to users, helping them specify their exact intent for improved conversation accuracy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intentDisambiguationSettingsProperty := &IntentDisambiguationSettingsProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	CustomDisambiguationMessage: jsii.String("customDisambiguationMessage"),
//   	MaxDisambiguationIntents: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intentdisambiguationsettings.html
//
type CfnBot_IntentDisambiguationSettingsProperty struct {
	// Determines whether the Intent Disambiguation feature is enabled.
	//
	// When set to `true` , Amazon Lex will present disambiguation options to users when multiple intents could match their input, with the default being `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intentdisambiguationsettings.html#cfn-lex-bot-intentdisambiguationsettings-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Provides a custom message that will be displayed before presenting the disambiguation options to users.
	//
	// This message helps set the context for users and can be customized to match your bot's tone and brand. If not specified, a default message will be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intentdisambiguationsettings.html#cfn-lex-bot-intentdisambiguationsettings-customdisambiguationmessage
	//
	CustomDisambiguationMessage *string `field:"optional" json:"customDisambiguationMessage" yaml:"customDisambiguationMessage"`
	// Specifies the maximum number of intent options (2-5) to present to users when disambiguation is needed.
	//
	// This setting determines how many intent options will be shown to users when the system detects ambiguous input. The default value is 3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intentdisambiguationsettings.html#cfn-lex-bot-intentdisambiguationsettings-maxdisambiguationintents
	//
	MaxDisambiguationIntents *float64 `field:"optional" json:"maxDisambiguationIntents" yaml:"maxDisambiguationIntents"`
}


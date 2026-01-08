package awslex


// Provides configuration information for a locale.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html
//
type CfnBot_BotLocaleProperty struct {
	// The identifier of the language and locale that the bot will be used in.
	//
	// The string must match one of the supported locales.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-localeid
	//
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
	// Determines the threshold where Amazon Lex will insert the `AMAZON.FallbackIntent` , `AMAZON.KendraSearchIntent` , or both when returning alternative intents. You must configure an `AMAZON.FallbackIntent` . `AMAZON.KendraSearchIntent` is only inserted if it is configured for the bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-nluconfidencethreshold
	//
	NluConfidenceThreshold *float64 `field:"required" json:"nluConfidenceThreshold" yaml:"nluConfidenceThreshold"`
	// Specifies a custom vocabulary to use with a specific locale.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-customvocabulary
	//
	CustomVocabulary interface{} `field:"optional" json:"customVocabulary" yaml:"customVocabulary"`
	// A description of the bot locale.
	//
	// Use this to help identify the bot locale in lists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-generativeaisettings
	//
	GenerativeAiSettings interface{} `field:"optional" json:"generativeAiSettings" yaml:"generativeAiSettings"`
	// One or more intents defined for the locale.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-intents
	//
	Intents interface{} `field:"optional" json:"intents" yaml:"intents"`
	// One or more slot types defined for the locale.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-slottypes
	//
	SlotTypes interface{} `field:"optional" json:"slotTypes" yaml:"slotTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-speechdetectionsensitivity
	//
	SpeechDetectionSensitivity *string `field:"optional" json:"speechDetectionSensitivity" yaml:"speechDetectionSensitivity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-unifiedspeechsettings
	//
	UnifiedSpeechSettings interface{} `field:"optional" json:"unifiedSpeechSettings" yaml:"unifiedSpeechSettings"`
	// Defines settings for using an Amazon Polly voice to communicate with a user.
	//
	// Valid values include:
	//
	// - `standard`
	// - `neural`
	// - `long-form`
	// - `generative`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-voicesettings
	//
	VoiceSettings interface{} `field:"optional" json:"voiceSettings" yaml:"voiceSettings"`
}


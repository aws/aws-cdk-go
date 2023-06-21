package awslex


// Provides configuration information for a locale.
//
// Example:
//
//
type CfnBot_BotLocaleProperty struct {
	// The identifier of the language and locale that the bot will be used in.
	//
	// The string must match one of the supported locales.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
	// Determines the threshold where Amazon Lex will insert the `AMAZON.FallbackIntent` , `AMAZON.KendraSearchIntent` , or both when returning alternative intents. You must configure an `AMAZON.FallbackIntent` . `AMAZON.KendraSearchIntent` is only inserted if it is configured for the bot.
	NluConfidenceThreshold *float64 `field:"required" json:"nluConfidenceThreshold" yaml:"nluConfidenceThreshold"`
	// Specifies a custom vocabulary to use with a specific locale.
	CustomVocabulary interface{} `field:"optional" json:"customVocabulary" yaml:"customVocabulary"`
	// A description of the bot locale.
	//
	// Use this to help identify the bot locale in lists.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// One or more intents defined for the locale.
	Intents interface{} `field:"optional" json:"intents" yaml:"intents"`
	// One or more slot types defined for the locale.
	SlotTypes interface{} `field:"optional" json:"slotTypes" yaml:"slotTypes"`
	// Defines settings for using an Amazon Polly voice to communicate with a user.
	VoiceSettings interface{} `field:"optional" json:"voiceSettings" yaml:"voiceSettings"`
}


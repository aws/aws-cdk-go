package awslex


// Properties for defining a `CfnBot`.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-bot.html
//
type CfnBotProps struct {
	// By default, data stored by Amazon Lex is encrypted.
	//
	// The `DataPrivacy` structure provides settings that determine how Amazon Lex handles special cases of securing the data for your bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-bot.html#cfn-lex-bot-dataprivacy
	//
	DataPrivacy interface{} `field:"required" json:"dataPrivacy" yaml:"dataPrivacy"`
	// The time, in seconds, that Amazon Lex should keep information about a user's conversation with the bot.
	//
	// A user interaction remains active for the amount of time specified. If no conversation occurs during this time, the session expires and Amazon Lex deletes any data provided before the timeout.
	//
	// You can specify between 60 (1 minute) and 86,400 (24 hours) seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-bot.html#cfn-lex-bot-idlesessionttlinseconds
	//
	IdleSessionTtlInSeconds *float64 `field:"required" json:"idleSessionTtlInSeconds" yaml:"idleSessionTtlInSeconds"`
	// The name of the bot locale.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-bot.html#cfn-lex-bot-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the IAM role used to build and run the bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-bot.html#cfn-lex-bot-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Indicates whether Amazon Lex V2 should automatically build the locales for the bot after a change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-bot.html#cfn-lex-bot-autobuildbotlocales
	//
	AutoBuildBotLocales interface{} `field:"optional" json:"autoBuildBotLocales" yaml:"autoBuildBotLocales"`
	// The Amazon S3 location of files used to import a bot.
	//
	// The files must be in the import format specified in [JSON format for importing and exporting](https://docs.aws.amazon.com/lexv2/latest/dg/import-export-format.html) in the *Amazon Lex developer guide.*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-bot.html#cfn-lex-bot-botfiles3location
	//
	BotFileS3Location interface{} `field:"optional" json:"botFileS3Location" yaml:"botFileS3Location"`
	// A list of locales for the bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-bot.html#cfn-lex-bot-botlocales
	//
	BotLocales interface{} `field:"optional" json:"botLocales" yaml:"botLocales"`
	// A list of tags to add to the bot.
	//
	// You can only add tags when you import a bot. You can't use the `UpdateBot` operation to update tags. To update tags, use the `TagResource` operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-bot.html#cfn-lex-bot-bottags
	//
	BotTags interface{} `field:"optional" json:"botTags" yaml:"botTags"`
	// The description of the version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-bot.html#cfn-lex-bot-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-bot.html#cfn-lex-bot-replication
	//
	Replication interface{} `field:"optional" json:"replication" yaml:"replication"`
	// Specifies configuration settings for the alias used to test the bot.
	//
	// If the `TestBotAliasSettings` property is not specified, the settings are configured with default values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-bot.html#cfn-lex-bot-testbotaliassettings
	//
	TestBotAliasSettings interface{} `field:"optional" json:"testBotAliasSettings" yaml:"testBotAliasSettings"`
	// A list of tags to add to the test alias for a bot.
	//
	// You can only add tags when you import a bot. You can't use the `UpdateAlias` operation to update tags. To update tags on the test alias, use the `TagResource` operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-bot.html#cfn-lex-bot-testbotaliastags
	//
	TestBotAliasTags interface{} `field:"optional" json:"testBotAliasTags" yaml:"testBotAliasTags"`
}


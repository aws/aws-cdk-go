package previewawslexmixins


// Specifies configuration settings for the alias used to test the bot.
//
// If the `TestBotAliasSettings` property is not specified, the settings are configured with default values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var sentimentAnalysisSettings interface{}
//
//   testBotAliasSettingsProperty := &TestBotAliasSettingsProperty{
//   	BotAliasLocaleSettings: []interface{}{
//   		&BotAliasLocaleSettingsItemProperty{
//   			BotAliasLocaleSetting: &BotAliasLocaleSettingsProperty{
//   				CodeHookSpecification: &CodeHookSpecificationProperty{
//   					LambdaCodeHook: &LambdaCodeHookProperty{
//   						CodeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   						LambdaArn: jsii.String("lambdaArn"),
//   					},
//   				},
//   				Enabled: jsii.Boolean(false),
//   			},
//   			LocaleId: jsii.String("localeId"),
//   		},
//   	},
//   	ConversationLogSettings: &ConversationLogSettingsProperty{
//   		AudioLogSettings: []interface{}{
//   			&AudioLogSettingProperty{
//   				Destination: &AudioLogDestinationProperty{
//   					S3Bucket: &S3BucketLogDestinationProperty{
//   						KmsKeyArn: jsii.String("kmsKeyArn"),
//   						LogPrefix: jsii.String("logPrefix"),
//   						S3BucketArn: jsii.String("s3BucketArn"),
//   					},
//   				},
//   				Enabled: jsii.Boolean(false),
//   			},
//   		},
//   		TextLogSettings: []interface{}{
//   			&TextLogSettingProperty{
//   				Destination: &TextLogDestinationProperty{
//   					CloudWatch: &CloudWatchLogGroupLogDestinationProperty{
//   						CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   						LogPrefix: jsii.String("logPrefix"),
//   					},
//   				},
//   				Enabled: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	SentimentAnalysisSettings: sentimentAnalysisSettings,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-testbotaliassettings.html
//
type CfnBotPropsMixin_TestBotAliasSettingsProperty struct {
	// Specifies settings that are unique to a locale.
	//
	// For example, you can use a different Lambda function depending on the bot's locale.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-testbotaliassettings.html#cfn-lex-bot-testbotaliassettings-botaliaslocalesettings
	//
	BotAliasLocaleSettings interface{} `field:"optional" json:"botAliasLocaleSettings" yaml:"botAliasLocaleSettings"`
	// Specifies settings for conversation logs that save audio, text, and metadata information for conversations with your users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-testbotaliassettings.html#cfn-lex-bot-testbotaliassettings-conversationlogsettings
	//
	ConversationLogSettings interface{} `field:"optional" json:"conversationLogSettings" yaml:"conversationLogSettings"`
	// Specifies a description for the test bot alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-testbotaliassettings.html#cfn-lex-bot-testbotaliassettings-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-testbotaliassettings.html#cfn-lex-bot-testbotaliassettings-sentimentanalysissettings
	//
	SentimentAnalysisSettings interface{} `field:"optional" json:"sentimentAnalysisSettings" yaml:"sentimentAnalysisSettings"`
}


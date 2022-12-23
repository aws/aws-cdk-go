package awslex


// Specifies configuration settings for the alias used to test the bot.
//
// If the `TestBotAliasSettings` property is not specified, the settings are configured with default values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sentimentAnalysisSettings interface{}
//
//   testBotAliasSettingsProperty := &testBotAliasSettingsProperty{
//   	botAliasLocaleSettings: []interface{}{
//   		&botAliasLocaleSettingsItemProperty{
//   			botAliasLocaleSetting: &botAliasLocaleSettingsProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				codeHookSpecification: &codeHookSpecificationProperty{
//   					lambdaCodeHook: &lambdaCodeHookProperty{
//   						codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   						lambdaArn: jsii.String("lambdaArn"),
//   					},
//   				},
//   			},
//   			localeId: jsii.String("localeId"),
//   		},
//   	},
//   	conversationLogSettings: &conversationLogSettingsProperty{
//   		audioLogSettings: []interface{}{
//   			&audioLogSettingProperty{
//   				destination: &audioLogDestinationProperty{
//   					s3Bucket: &s3BucketLogDestinationProperty{
//   						logPrefix: jsii.String("logPrefix"),
//   						s3BucketArn: jsii.String("s3BucketArn"),
//
//   						// the properties below are optional
//   						kmsKeyArn: jsii.String("kmsKeyArn"),
//   					},
//   				},
//   				enabled: jsii.Boolean(false),
//   			},
//   		},
//   		textLogSettings: []interface{}{
//   			&textLogSettingProperty{
//   				destination: &textLogDestinationProperty{
//   					cloudWatch: &cloudWatchLogGroupLogDestinationProperty{
//   						cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   						logPrefix: jsii.String("logPrefix"),
//   					},
//   				},
//   				enabled: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	sentimentAnalysisSettings: sentimentAnalysisSettings,
//   }
//
type CfnBot_TestBotAliasSettingsProperty struct {
	// Specifies settings that are unique to a locale.
	//
	// For example, you can use a different Lambda function depending on the bot's locale.
	BotAliasLocaleSettings interface{} `field:"optional" json:"botAliasLocaleSettings" yaml:"botAliasLocaleSettings"`
	// Specifies settings for conversation logs that save audio, text, and metadata information for conversations with your users.
	ConversationLogSettings interface{} `field:"optional" json:"conversationLogSettings" yaml:"conversationLogSettings"`
	// Specifies a description for the test bot alias.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
	SentimentAnalysisSettings interface{} `field:"optional" json:"sentimentAnalysisSettings" yaml:"sentimentAnalysisSettings"`
}


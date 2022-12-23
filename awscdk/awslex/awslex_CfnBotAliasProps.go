package awslex


// Properties for defining a `CfnBotAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sentimentAnalysisSettings interface{}
//
//   cfnBotAliasProps := &cfnBotAliasProps{
//   	botAliasName: jsii.String("botAliasName"),
//   	botId: jsii.String("botId"),
//
//   	// the properties below are optional
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
//   	botAliasTags: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	botVersion: jsii.String("botVersion"),
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
type CfnBotAliasProps struct {
	// The name of the bot alias.
	BotAliasName *string `field:"required" json:"botAliasName" yaml:"botAliasName"`
	// The unique identifier of the bot.
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// Maps configuration information to a specific locale.
	//
	// You can use this parameter to specify a specific Lambda function to run different functions in different locales.
	BotAliasLocaleSettings interface{} `field:"optional" json:"botAliasLocaleSettings" yaml:"botAliasLocaleSettings"`
	// An array of key-value pairs to apply to this resource.
	//
	// You can only add tags when you specify an alias.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	BotAliasTags interface{} `field:"optional" json:"botAliasTags" yaml:"botAliasTags"`
	// The version of the bot that the bot alias references.
	BotVersion *string `field:"optional" json:"botVersion" yaml:"botVersion"`
	// Specifies whether Amazon Lex logs text and audio for conversations with the bot.
	//
	// When you enable conversation logs, text logs store text input, transcripts of audio input, and associated metadata in Amazon CloudWatch logs. Audio logs store input in Amazon S3 .
	ConversationLogSettings interface{} `field:"optional" json:"conversationLogSettings" yaml:"conversationLogSettings"`
	// The description of the bot alias.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
	SentimentAnalysisSettings interface{} `field:"optional" json:"sentimentAnalysisSettings" yaml:"sentimentAnalysisSettings"`
}


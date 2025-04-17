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
//   cfnBotAliasProps := &CfnBotAliasProps{
//   	BotAliasName: jsii.String("botAliasName"),
//   	BotId: jsii.String("botId"),
//
//   	// the properties below are optional
//   	BotAliasLocaleSettings: []interface{}{
//   		&BotAliasLocaleSettingsItemProperty{
//   			BotAliasLocaleSetting: &BotAliasLocaleSettingsProperty{
//   				Enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				CodeHookSpecification: &CodeHookSpecificationProperty{
//   					LambdaCodeHook: &LambdaCodeHookProperty{
//   						CodeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   						LambdaArn: jsii.String("lambdaArn"),
//   					},
//   				},
//   			},
//   			LocaleId: jsii.String("localeId"),
//   		},
//   	},
//   	BotAliasTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	BotVersion: jsii.String("botVersion"),
//   	ConversationLogSettings: &ConversationLogSettingsProperty{
//   		AudioLogSettings: []interface{}{
//   			&AudioLogSettingProperty{
//   				Destination: &AudioLogDestinationProperty{
//   					S3Bucket: &S3BucketLogDestinationProperty{
//   						LogPrefix: jsii.String("logPrefix"),
//   						S3BucketArn: jsii.String("s3BucketArn"),
//
//   						// the properties below are optional
//   						KmsKeyArn: jsii.String("kmsKeyArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botalias.html
//
type CfnBotAliasProps struct {
	// The name of the bot alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botalias.html#cfn-lex-botalias-botaliasname
	//
	BotAliasName *string `field:"required" json:"botAliasName" yaml:"botAliasName"`
	// The unique identifier of the bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botalias.html#cfn-lex-botalias-botid
	//
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// Specifies settings that are unique to a locale.
	//
	// For example, you can use different Lambda function depending on the bot's locale.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botalias.html#cfn-lex-botalias-botaliaslocalesettings
	//
	BotAliasLocaleSettings interface{} `field:"optional" json:"botAliasLocaleSettings" yaml:"botAliasLocaleSettings"`
	// An array of key-value pairs to apply to this resource.
	//
	// You can only add tags when you specify an alias.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botalias.html#cfn-lex-botalias-botaliastags
	//
	BotAliasTags interface{} `field:"optional" json:"botAliasTags" yaml:"botAliasTags"`
	// The version of the bot that the bot alias references.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botalias.html#cfn-lex-botalias-botversion
	//
	BotVersion *string `field:"optional" json:"botVersion" yaml:"botVersion"`
	// Specifies whether Amazon Lex logs text and audio for conversations with the bot.
	//
	// When you enable conversation logs, text logs store text input, transcripts of audio input, and associated metadata in Amazon CloudWatch logs. Audio logs store input in Amazon S3 .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botalias.html#cfn-lex-botalias-conversationlogsettings
	//
	ConversationLogSettings interface{} `field:"optional" json:"conversationLogSettings" yaml:"conversationLogSettings"`
	// The description of the bot alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botalias.html#cfn-lex-botalias-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botalias.html#cfn-lex-botalias-sentimentanalysissettings
	//
	SentimentAnalysisSettings interface{} `field:"optional" json:"sentimentAnalysisSettings" yaml:"sentimentAnalysisSettings"`
}


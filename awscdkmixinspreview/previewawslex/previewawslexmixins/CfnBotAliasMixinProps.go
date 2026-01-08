package previewawslexmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnBotAliasPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var sentimentAnalysisSettings interface{}
//
//   cfnBotAliasMixinProps := &CfnBotAliasMixinProps{
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
//   	BotAliasName: jsii.String("botAliasName"),
//   	BotAliasTags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	BotId: jsii.String("botId"),
//   	BotVersion: jsii.String("botVersion"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botalias.html
//
type CfnBotAliasMixinProps struct {
	// Specifies settings that are unique to a locale.
	//
	// For example, you can use different Lambda function depending on the bot's locale.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botalias.html#cfn-lex-botalias-botaliaslocalesettings
	//
	BotAliasLocaleSettings interface{} `field:"optional" json:"botAliasLocaleSettings" yaml:"botAliasLocaleSettings"`
	// The name of the bot alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botalias.html#cfn-lex-botalias-botaliasname
	//
	BotAliasName *string `field:"optional" json:"botAliasName" yaml:"botAliasName"`
	// An array of key-value pairs to apply to this resource.
	//
	// You can only add tags when you specify an alias.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botalias.html#cfn-lex-botalias-botaliastags
	//
	BotAliasTags *[]*awscdk.CfnTag `field:"optional" json:"botAliasTags" yaml:"botAliasTags"`
	// The unique identifier of the bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botalias.html#cfn-lex-botalias-botid
	//
	BotId *string `field:"optional" json:"botId" yaml:"botId"`
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


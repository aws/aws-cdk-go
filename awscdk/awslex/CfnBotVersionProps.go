package awslex


// Properties for defining a `CfnBotVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBotVersionProps := &CfnBotVersionProps{
//   	BotId: jsii.String("botId"),
//   	BotVersionLocaleSpecification: []interface{}{
//   		&BotVersionLocaleSpecificationProperty{
//   			BotVersionLocaleDetails: &BotVersionLocaleDetailsProperty{
//   				SourceBotVersion: jsii.String("sourceBotVersion"),
//   			},
//   			LocaleId: jsii.String("localeId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botversion.html
//
type CfnBotVersionProps struct {
	// The unique identifier of the bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botversion.html#cfn-lex-botversion-botid
	//
	BotId interface{} `field:"required" json:"botId" yaml:"botId"`
	// Specifies the locales that Amazon Lex adds to this version.
	//
	// You can choose the Draft version or any other previously published version for each locale. When you specify a source version, the locale data is copied from the source version to the new version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botversion.html#cfn-lex-botversion-botversionlocalespecification
	//
	BotVersionLocaleSpecification interface{} `field:"required" json:"botVersionLocaleSpecification" yaml:"botVersionLocaleSpecification"`
	// The description of the version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botversion.html#cfn-lex-botversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}


package awslex


// Properties for defining a `CfnBotVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBotVersionProps := &cfnBotVersionProps{
//   	botId: jsii.String("botId"),
//   	botVersionLocaleSpecification: []interface{}{
//   		&botVersionLocaleSpecificationProperty{
//   			botVersionLocaleDetails: &botVersionLocaleDetailsProperty{
//   				sourceBotVersion: jsii.String("sourceBotVersion"),
//   			},
//   			localeId: jsii.String("localeId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnBotVersionProps struct {
	// The unique identifier of the bot.
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// Specifies the locales that Amazon Lex adds to this version.
	//
	// You can choose the Draft version or any other previously published version for each locale. When you specify a source version, the locale data is copied from the source version to the new version.
	BotVersionLocaleSpecification interface{} `field:"required" json:"botVersionLocaleSpecification" yaml:"botVersionLocaleSpecification"`
	// The description of the version.
	Description *string `field:"optional" json:"description" yaml:"description"`
}


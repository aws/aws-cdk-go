package awslex


// Specifies the locale that Amazon Lex adds to this version.
//
// You can choose the Draft version or any other previously published version for each locale. When you specify a source version, the locale data is copied from the source version to the new version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botVersionLocaleSpecificationProperty := &botVersionLocaleSpecificationProperty{
//   	botVersionLocaleDetails: &botVersionLocaleDetailsProperty{
//   		sourceBotVersion: jsii.String("sourceBotVersion"),
//   	},
//   	localeId: jsii.String("localeId"),
//   }
//
type CfnBotVersion_BotVersionLocaleSpecificationProperty struct {
	// The version of a bot used for a bot locale.
	BotVersionLocaleDetails interface{} `field:"required" json:"botVersionLocaleDetails" yaml:"botVersionLocaleDetails"`
	// The identifier of the locale to add to the version.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
}


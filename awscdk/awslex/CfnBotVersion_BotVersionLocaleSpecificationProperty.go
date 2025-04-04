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
//   botVersionLocaleSpecificationProperty := &BotVersionLocaleSpecificationProperty{
//   	BotVersionLocaleDetails: &BotVersionLocaleDetailsProperty{
//   		SourceBotVersion: jsii.String("sourceBotVersion"),
//   	},
//   	LocaleId: jsii.String("localeId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botversion-botversionlocalespecification.html
//
type CfnBotVersion_BotVersionLocaleSpecificationProperty struct {
	// The version of a bot used for a bot locale.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botversion-botversionlocalespecification.html#cfn-lex-botversion-botversionlocalespecification-botversionlocaledetails
	//
	BotVersionLocaleDetails interface{} `field:"required" json:"botVersionLocaleDetails" yaml:"botVersionLocaleDetails"`
	// The identifier of the locale to add to the version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botversion-botversionlocalespecification.html#cfn-lex-botversion-botversionlocalespecification-localeid
	//
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
}


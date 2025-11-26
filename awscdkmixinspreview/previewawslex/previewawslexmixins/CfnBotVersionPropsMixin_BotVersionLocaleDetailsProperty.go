package previewawslexmixins


// The version of a bot used for a bot locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   botVersionLocaleDetailsProperty := &BotVersionLocaleDetailsProperty{
//   	SourceBotVersion: jsii.String("sourceBotVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botversion-botversionlocaledetails.html
//
type CfnBotVersionPropsMixin_BotVersionLocaleDetailsProperty struct {
	// The version of a bot used for a bot locale.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botversion-botversionlocaledetails.html#cfn-lex-botversion-botversionlocaledetails-sourcebotversion
	//
	SourceBotVersion *string `field:"optional" json:"sourceBotVersion" yaml:"sourceBotVersion"`
}


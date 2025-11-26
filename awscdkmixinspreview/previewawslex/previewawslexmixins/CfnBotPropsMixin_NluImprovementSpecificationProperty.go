package previewawslexmixins


// Configures the Assisted Natural Language Understanding (NLU) feature for your bot.
//
// This specification determines whether enhanced intent recognition and utterance understanding capabilities are active.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   nluImprovementSpecificationProperty := &NluImprovementSpecificationProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-nluimprovementspecification.html
//
type CfnBotPropsMixin_NluImprovementSpecificationProperty struct {
	// Determines whether the Assisted NLU feature is enabled for the bot.
	//
	// When set to `true` , Amazon Lex uses advanced models to improve intent recognition and slot resolution, with the default being `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-nluimprovementspecification.html#cfn-lex-bot-nluimprovementspecification-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}


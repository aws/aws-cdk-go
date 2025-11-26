package previewawslexmixins


// Determines whether Amazon Lex obscures slot values in conversation logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   obfuscationSettingProperty := &ObfuscationSettingProperty{
//   	ObfuscationSettingType: jsii.String("obfuscationSettingType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-obfuscationsetting.html
//
type CfnBotPropsMixin_ObfuscationSettingProperty struct {
	// Value that determines whether Amazon Lex obscures slot values in conversation logs.
	//
	// The default is to obscure the values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-obfuscationsetting.html#cfn-lex-bot-obfuscationsetting-obfuscationsettingtype
	//
	ObfuscationSettingType *string `field:"optional" json:"obfuscationSettingType" yaml:"obfuscationSettingType"`
}


package awslex


// Determines whether Amazon Lex obscures slot values in conversation logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   obfuscationSettingProperty := &ObfuscationSettingProperty{
//   	ObfuscationSettingType: jsii.String("obfuscationSettingType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-obfuscationsetting.html
//
type CfnBot_ObfuscationSettingProperty struct {
	// Value that determines whether Amazon Lex obscures slot values in conversation logs.
	//
	// The default is to obscure the values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-obfuscationsetting.html#cfn-lex-bot-obfuscationsetting-obfuscationsettingtype
	//
	ObfuscationSettingType *string `field:"required" json:"obfuscationSettingType" yaml:"obfuscationSettingType"`
}


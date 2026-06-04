package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   audioFillerSettingsProperty := &AudioFillerSettingsProperty{
//   	AudioType: jsii.String("audioType"),
//   	Enabled: jsii.Boolean(false),
//   	MinimumPlayDurationInMilliseconds: jsii.Number(123),
//   	ResponseDeliveryDelayInMilliseconds: jsii.Number(123),
//   	StartDelayInMilliseconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiofillersettings.html
//
type CfnBotPropsMixin_AudioFillerSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiofillersettings.html#cfn-lex-bot-audiofillersettings-audiotype
	//
	AudioType *string `field:"optional" json:"audioType" yaml:"audioType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiofillersettings.html#cfn-lex-bot-audiofillersettings-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiofillersettings.html#cfn-lex-bot-audiofillersettings-minimumplaydurationinmilliseconds
	//
	MinimumPlayDurationInMilliseconds *float64 `field:"optional" json:"minimumPlayDurationInMilliseconds" yaml:"minimumPlayDurationInMilliseconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiofillersettings.html#cfn-lex-bot-audiofillersettings-responsedeliverydelayinmilliseconds
	//
	ResponseDeliveryDelayInMilliseconds *float64 `field:"optional" json:"responseDeliveryDelayInMilliseconds" yaml:"responseDeliveryDelayInMilliseconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiofillersettings.html#cfn-lex-bot-audiofillersettings-startdelayinmilliseconds
	//
	StartDelayInMilliseconds *float64 `field:"optional" json:"startDelayInMilliseconds" yaml:"startDelayInMilliseconds"`
}


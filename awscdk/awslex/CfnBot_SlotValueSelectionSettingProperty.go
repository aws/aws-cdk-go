package awslex


// Contains settings used by Amazon Lex to select a slot value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotValueSelectionSettingProperty := &SlotValueSelectionSettingProperty{
//   	ResolutionStrategy: jsii.String("resolutionStrategy"),
//
//   	// the properties below are optional
//   	AdvancedRecognitionSetting: &AdvancedRecognitionSettingProperty{
//   		AudioRecognitionStrategy: jsii.String("audioRecognitionStrategy"),
//   	},
//   	RegexFilter: &SlotValueRegexFilterProperty{
//   		Pattern: jsii.String("pattern"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotvalueselectionsetting.html
//
type CfnBot_SlotValueSelectionSettingProperty struct {
	// Determines the slot resolution strategy that Amazon Lex uses to return slot type values.
	//
	// The field can be set to one of the following values:
	//
	// - `ORIGINAL_VALUE` - Returns the value entered by the user, if the user value is similar to the slot value.
	// - `TOP_RESOLUTION` - If there is a resolution list for the slot, return the first value in the resolution list as the slot type value. If there is no resolution list, null is returned.
	//
	// If you don't specify the `valueSelectionStrategy` , the default is `ORIGINAL_VALUE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotvalueselectionsetting.html#cfn-lex-bot-slotvalueselectionsetting-resolutionstrategy
	//
	ResolutionStrategy *string `field:"required" json:"resolutionStrategy" yaml:"resolutionStrategy"`
	// Provides settings that enable advanced recognition settings for slot values.
	//
	// You can use this to enable using slot values as a custom vocabulary for recognizing user utterances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotvalueselectionsetting.html#cfn-lex-bot-slotvalueselectionsetting-advancedrecognitionsetting
	//
	AdvancedRecognitionSetting interface{} `field:"optional" json:"advancedRecognitionSetting" yaml:"advancedRecognitionSetting"`
	// A regular expression used to validate the value of a slot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotvalueselectionsetting.html#cfn-lex-bot-slotvalueselectionsetting-regexfilter
	//
	RegexFilter interface{} `field:"optional" json:"regexFilter" yaml:"regexFilter"`
}


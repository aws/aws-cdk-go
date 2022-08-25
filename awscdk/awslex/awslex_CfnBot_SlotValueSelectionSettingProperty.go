package awslex


// Contains settings used by Amazon Lex to select a slot value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotValueSelectionSettingProperty := &slotValueSelectionSettingProperty{
//   	resolutionStrategy: jsii.String("resolutionStrategy"),
//
//   	// the properties below are optional
//   	advancedRecognitionSetting: &advancedRecognitionSettingProperty{
//   		audioRecognitionStrategy: jsii.String("audioRecognitionStrategy"),
//   	},
//   	regexFilter: &slotValueRegexFilterProperty{
//   		pattern: jsii.String("pattern"),
//   	},
//   }
//
type CfnBot_SlotValueSelectionSettingProperty struct {
	// Determines the slot resolution strategy that Amazon Lex uses to return slot type values.
	//
	// The field can be set to one of the following values:
	//
	// - OriginalValue - Returns the value entered by the user, if the user value is similar to a slot value.
	// - TopResolution - If there is a resolution list for the slot, return the first value in the resolution list as the slot type value. If there is no resolution list, null is returned.
	//
	// If you don't specify the valueSelectionStrategy, the default is OriginalValue.
	ResolutionStrategy *string `field:"required" json:"resolutionStrategy" yaml:"resolutionStrategy"`
	// Specifies settings that enable advanced recognition settings for slot values.
	//
	// You can use this to enable using slot values as a custom vocabulary for recognizing user utterances.
	AdvancedRecognitionSetting interface{} `field:"optional" json:"advancedRecognitionSetting" yaml:"advancedRecognitionSetting"`
	// A regular expression used to validate the value of a slot.
	RegexFilter interface{} `field:"optional" json:"regexFilter" yaml:"regexFilter"`
}


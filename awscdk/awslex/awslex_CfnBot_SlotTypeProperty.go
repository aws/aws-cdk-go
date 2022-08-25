package awslex


// Describes a slot type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotTypeProperty := &slotTypeProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	externalSourceSetting: &externalSourceSettingProperty{
//   		grammarSlotTypeSetting: &grammarSlotTypeSettingProperty{
//   			source: &grammarSlotTypeSourceProperty{
//   				s3BucketName: jsii.String("s3BucketName"),
//   				s3ObjectKey: jsii.String("s3ObjectKey"),
//
//   				// the properties below are optional
//   				kmsKeyArn: jsii.String("kmsKeyArn"),
//   			},
//   		},
//   	},
//   	parentSlotTypeSignature: jsii.String("parentSlotTypeSignature"),
//   	slotTypeValues: []interface{}{
//   		&slotTypeValueProperty{
//   			sampleValue: &sampleValueProperty{
//   				value: jsii.String("value"),
//   			},
//
//   			// the properties below are optional
//   			synonyms: []interface{}{
//   				&sampleValueProperty{
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	valueSelectionSetting: &slotValueSelectionSettingProperty{
//   		resolutionStrategy: jsii.String("resolutionStrategy"),
//
//   		// the properties below are optional
//   		advancedRecognitionSetting: &advancedRecognitionSettingProperty{
//   			audioRecognitionStrategy: jsii.String("audioRecognitionStrategy"),
//   		},
//   		regexFilter: &slotValueRegexFilterProperty{
//   			pattern: jsii.String("pattern"),
//   		},
//   	},
//   }
//
type CfnBot_SlotTypeProperty struct {
	// The name of the slot type.
	//
	// A slot type name must be unique withing the account.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the slot type.
	//
	// Use the description to help identify the slot type in lists.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Sets the type of external information used to create the slot type.
	ExternalSourceSetting interface{} `field:"optional" json:"externalSourceSetting" yaml:"externalSourceSetting"`
	// The built-in slot type used as a parent of this slot type.
	//
	// When you define a parent slot type, the new slot type has the configuration of the parent lot type.
	//
	// Only AMAZON.AlphaNumeric is supported.
	ParentSlotTypeSignature *string `field:"optional" json:"parentSlotTypeSignature" yaml:"parentSlotTypeSignature"`
	// A list of SlotTypeValue objects that defines the values that the slot type can take.
	//
	// Each value can have a list of synonyms, additional values that help train the machine learning model about the values that it resolves for the slot.
	SlotTypeValues interface{} `field:"optional" json:"slotTypeValues" yaml:"slotTypeValues"`
	// Determines the slot resolution strategy that Amazon Lex uses to return slot type values.
	//
	// The field can be set to one of the following values:
	//
	// - OriginalValue - Returns the value entered by the user, if the user value is similar to a slot value.
	// - TopResolution - If there is a resolution list for the slot, return the first value in the resolution list as the slot type value. If there is no resolution list, null is returned.
	//
	// If you don't specify the valueSelectionStrategy, the default is OriginalValue.
	ValueSelectionSetting interface{} `field:"optional" json:"valueSelectionSetting" yaml:"valueSelectionSetting"`
}


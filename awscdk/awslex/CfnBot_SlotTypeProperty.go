package awslex


// Describes a slot type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotTypeProperty := &SlotTypeProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CompositeSlotTypeSetting: &CompositeSlotTypeSettingProperty{
//   		SubSlots: []interface{}{
//   			&SubSlotTypeCompositionProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				SlotTypeId: jsii.String("slotTypeId"),
//   				SlotTypeName: jsii.String("slotTypeName"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ExternalSourceSetting: &ExternalSourceSettingProperty{
//   		GrammarSlotTypeSetting: &GrammarSlotTypeSettingProperty{
//   			Source: &GrammarSlotTypeSourceProperty{
//   				S3BucketName: jsii.String("s3BucketName"),
//   				S3ObjectKey: jsii.String("s3ObjectKey"),
//
//   				// the properties below are optional
//   				KmsKeyArn: jsii.String("kmsKeyArn"),
//   			},
//   		},
//   	},
//   	ParentSlotTypeSignature: jsii.String("parentSlotTypeSignature"),
//   	SlotTypeValues: []interface{}{
//   		&SlotTypeValueProperty{
//   			SampleValue: &SampleValueProperty{
//   				Value: jsii.String("value"),
//   			},
//
//   			// the properties below are optional
//   			Synonyms: []interface{}{
//   				&SampleValueProperty{
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	ValueSelectionSetting: &SlotValueSelectionSettingProperty{
//   		ResolutionStrategy: jsii.String("resolutionStrategy"),
//
//   		// the properties below are optional
//   		AdvancedRecognitionSetting: &AdvancedRecognitionSettingProperty{
//   			AudioRecognitionStrategy: jsii.String("audioRecognitionStrategy"),
//   		},
//   		RegexFilter: &SlotValueRegexFilterProperty{
//   			Pattern: jsii.String("pattern"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slottype.html
//
type CfnBot_SlotTypeProperty struct {
	// The name of the slot type.
	//
	// A slot type name must be unique withing the account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slottype.html#cfn-lex-bot-slottype-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slottype.html#cfn-lex-bot-slottype-compositeslottypesetting
	//
	CompositeSlotTypeSetting interface{} `field:"optional" json:"compositeSlotTypeSetting" yaml:"compositeSlotTypeSetting"`
	// A description of the slot type.
	//
	// Use the description to help identify the slot type in lists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slottype.html#cfn-lex-bot-slottype-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Sets the type of external information used to create the slot type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slottype.html#cfn-lex-bot-slottype-externalsourcesetting
	//
	ExternalSourceSetting interface{} `field:"optional" json:"externalSourceSetting" yaml:"externalSourceSetting"`
	// The built-in slot type used as a parent of this slot type.
	//
	// When you define a parent slot type, the new slot type has the configuration of the parent lot type.
	//
	// Only `AMAZON.AlphaNumeric` is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slottype.html#cfn-lex-bot-slottype-parentslottypesignature
	//
	ParentSlotTypeSignature *string `field:"optional" json:"parentSlotTypeSignature" yaml:"parentSlotTypeSignature"`
	// A list of SlotTypeValue objects that defines the values that the slot type can take.
	//
	// Each value can have a list of synonyms, additional values that help train the machine learning model about the values that it resolves for the slot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slottype.html#cfn-lex-bot-slottype-slottypevalues
	//
	SlotTypeValues interface{} `field:"optional" json:"slotTypeValues" yaml:"slotTypeValues"`
	// Determines the slot resolution strategy that Amazon Lex uses to return slot type values.
	//
	// The field can be set to one of the following values:
	//
	// - `ORIGINAL_VALUE` - Returns the value entered by the user, if the user value is similar to the slot value.
	// - `TOP_RESOLUTION` - If there is a resolution list for the slot, return the first value in the resolution list as the slot type value. If there is no resolution list, null is returned.
	//
	// If you don't specify the `valueSelectionStrategy` , the default is `ORIGINAL_VALUE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slottype.html#cfn-lex-bot-slottype-valueselectionsetting
	//
	ValueSelectionSetting interface{} `field:"optional" json:"valueSelectionSetting" yaml:"valueSelectionSetting"`
}


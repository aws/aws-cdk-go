package awsbedrock


// Settings for generating data from audio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioExtractionCategoryProperty := &AudioExtractionCategoryProperty{
//   	State: jsii.String("state"),
//
//   	// the properties below are optional
//   	TypeConfiguration: &AudioExtractionCategoryTypeConfigurationProperty{
//   		Transcript: &TranscriptConfigurationProperty{
//   			ChannelLabeling: &ChannelLabelingConfigurationProperty{
//   				State: jsii.String("state"),
//   			},
//   			SpeakerLabeling: &SpeakerLabelingConfigurationProperty{
//   				State: jsii.String("state"),
//   			},
//   		},
//   	},
//   	Types: []*string{
//   		jsii.String("types"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audioextractioncategory.html
//
type CfnDataAutomationProject_AudioExtractionCategoryProperty struct {
	// Whether generating categorical data from audio is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audioextractioncategory.html#cfn-bedrock-dataautomationproject-audioextractioncategory-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
	// This element contains information about extractions from different types.
	//
	// Used to enable speaker and channel labeling for transcripts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audioextractioncategory.html#cfn-bedrock-dataautomationproject-audioextractioncategory-typeconfiguration
	//
	TypeConfiguration interface{} `field:"optional" json:"typeConfiguration" yaml:"typeConfiguration"`
	// The types of data to generate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audioextractioncategory.html#cfn-bedrock-dataautomationproject-audioextractioncategory-types
	//
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}


package awsbedrock


// Configuration settings for processing audio content in multimodal knowledge bases.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioConfigurationProperty := &AudioConfigurationProperty{
//   	SegmentationConfiguration: &AudioSegmentationConfigurationProperty{
//   		FixedLengthDuration: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-audioconfiguration.html
//
type CfnKnowledgeBase_AudioConfigurationProperty struct {
	// Configuration for segmenting audio content during processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-audioconfiguration.html#cfn-bedrock-knowledgebase-audioconfiguration-segmentationconfiguration
	//
	SegmentationConfiguration interface{} `field:"required" json:"segmentationConfiguration" yaml:"segmentationConfiguration"`
}


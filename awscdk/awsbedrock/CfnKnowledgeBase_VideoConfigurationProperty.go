package awsbedrock


// Configuration settings for processing video content in multimodal knowledge bases.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoConfigurationProperty := &VideoConfigurationProperty{
//   	SegmentationConfiguration: &VideoSegmentationConfigurationProperty{
//   		FixedLengthDuration: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-videoconfiguration.html
//
type CfnKnowledgeBase_VideoConfigurationProperty struct {
	// Configuration for segmenting video content during processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-videoconfiguration.html#cfn-bedrock-knowledgebase-videoconfiguration-segmentationconfiguration
	//
	SegmentationConfiguration interface{} `field:"required" json:"segmentationConfiguration" yaml:"segmentationConfiguration"`
}


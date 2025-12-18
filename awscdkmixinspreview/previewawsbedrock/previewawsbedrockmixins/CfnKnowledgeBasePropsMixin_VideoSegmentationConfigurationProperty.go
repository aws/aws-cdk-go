package previewawsbedrockmixins


// Configuration for segmenting video content during multimodal knowledge base ingestion.
//
// Determines how video files are divided into chunks for processing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   videoSegmentationConfigurationProperty := &VideoSegmentationConfigurationProperty{
//   	FixedLengthDuration: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-videosegmentationconfiguration.html
//
type CfnKnowledgeBasePropsMixin_VideoSegmentationConfigurationProperty struct {
	// The duration in seconds for each video segment.
	//
	// Video files will be divided into chunks of this length for processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-videosegmentationconfiguration.html#cfn-bedrock-knowledgebase-videosegmentationconfiguration-fixedlengthduration
	//
	FixedLengthDuration *float64 `field:"optional" json:"fixedLengthDuration" yaml:"fixedLengthDuration"`
}


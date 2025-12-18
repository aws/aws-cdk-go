package awsbedrock


// Configuration for segmenting audio content during multimodal knowledge base ingestion.
//
// Determines how audio files are divided into chunks for processing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioSegmentationConfigurationProperty := &AudioSegmentationConfigurationProperty{
//   	FixedLengthDuration: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-audiosegmentationconfiguration.html
//
type CfnKnowledgeBase_AudioSegmentationConfigurationProperty struct {
	// The duration in seconds for each audio segment.
	//
	// Audio files will be divided into chunks of this length for processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-audiosegmentationconfiguration.html#cfn-bedrock-knowledgebase-audiosegmentationconfiguration-fixedlengthduration
	//
	FixedLengthDuration *float64 `field:"required" json:"fixedLengthDuration" yaml:"fixedLengthDuration"`
}


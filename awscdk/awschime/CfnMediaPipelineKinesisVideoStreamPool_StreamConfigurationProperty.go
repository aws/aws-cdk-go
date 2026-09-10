package awschime


// The configuration settings for the Kinesis video stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamConfigurationProperty := &StreamConfigurationProperty{
//   	Region: jsii.String("region"),
//
//   	// the properties below are optional
//   	DataRetentionInHours: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediapipelinekinesisvideostreampool-streamconfiguration.html
//
type CfnMediaPipelineKinesisVideoStreamPool_StreamConfigurationProperty struct {
	// The AWS Region of the video stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediapipelinekinesisvideostreampool-streamconfiguration.html#cfn-chime-mediapipelinekinesisvideostreampool-streamconfiguration-region
	//
	Region *string `field:"required" json:"region" yaml:"region"`
	// The amount of time that data is retained, in hours.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediapipelinekinesisvideostreampool-streamconfiguration.html#cfn-chime-mediapipelinekinesisvideostreampool-streamconfiguration-dataretentioninhours
	//
	DataRetentionInHours *float64 `field:"optional" json:"dataRetentionInHours" yaml:"dataRetentionInHours"`
}


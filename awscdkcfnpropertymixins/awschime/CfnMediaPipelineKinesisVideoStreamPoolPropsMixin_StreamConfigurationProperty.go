package awschime


// The configuration settings for the Kinesis video stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   streamConfigurationProperty := &StreamConfigurationProperty{
//   	DataRetentionInHours: jsii.Number(123),
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediapipelinekinesisvideostreampool-streamconfiguration.html
//
type CfnMediaPipelineKinesisVideoStreamPoolPropsMixin_StreamConfigurationProperty struct {
	// The amount of time that data is retained, in hours.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediapipelinekinesisvideostreampool-streamconfiguration.html#cfn-chime-mediapipelinekinesisvideostreampool-streamconfiguration-dataretentioninhours
	//
	DataRetentionInHours *float64 `field:"optional" json:"dataRetentionInHours" yaml:"dataRetentionInHours"`
	// The AWS Region of the video stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediapipelinekinesisvideostreampool-streamconfiguration.html#cfn-chime-mediapipelinekinesisvideostreampool-streamconfiguration-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}


package awschime


// Properties for defining a `CfnMediaPipelineKinesisVideoStreamPool`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMediaPipelineKinesisVideoStreamPoolProps := &CfnMediaPipelineKinesisVideoStreamPoolProps{
//   	PoolName: jsii.String("poolName"),
//   	StreamConfiguration: &StreamConfigurationProperty{
//   		Region: jsii.String("region"),
//
//   		// the properties below are optional
//   		DataRetentionInHours: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediapipelinekinesisvideostreampool.html
//
type CfnMediaPipelineKinesisVideoStreamPoolProps struct {
	// The name of the Kinesis Video Stream Pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediapipelinekinesisvideostreampool.html#cfn-chime-mediapipelinekinesisvideostreampool-poolname
	//
	PoolName *string `field:"required" json:"poolName" yaml:"poolName"`
	// The configuration settings for the Kinesis video stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediapipelinekinesisvideostreampool.html#cfn-chime-mediapipelinekinesisvideostreampool-streamconfiguration
	//
	StreamConfiguration interface{} `field:"required" json:"streamConfiguration" yaml:"streamConfiguration"`
	// The tags associated with the Kinesis Video Stream Pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediapipelinekinesisvideostreampool.html#cfn-chime-mediapipelinekinesisvideostreampool-tags
	//
	Tags *[]*CfnMediaPipelineKinesisVideoStreamPool_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}


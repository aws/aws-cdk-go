package awschime


// Properties for CfnMediaPipelineKinesisVideoStreamPoolPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnMediaPipelineKinesisVideoStreamPoolMixinProps := &CfnMediaPipelineKinesisVideoStreamPoolMixinProps{
//   	PoolName: jsii.String("poolName"),
//   	StreamConfiguration: &StreamConfigurationProperty{
//   		DataRetentionInHours: jsii.Number(123),
//   		Region: jsii.String("region"),
//   	},
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
type CfnMediaPipelineKinesisVideoStreamPoolMixinProps struct {
	// The name of the Kinesis Video Stream Pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediapipelinekinesisvideostreampool.html#cfn-chime-mediapipelinekinesisvideostreampool-poolname
	//
	PoolName *string `field:"optional" json:"poolName" yaml:"poolName"`
	// The configuration settings for the Kinesis video stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediapipelinekinesisvideostreampool.html#cfn-chime-mediapipelinekinesisvideostreampool-streamconfiguration
	//
	StreamConfiguration interface{} `field:"optional" json:"streamConfiguration" yaml:"streamConfiguration"`
	// The tags associated with the Kinesis Video Stream Pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediapipelinekinesisvideostreampool.html#cfn-chime-mediapipelinekinesisvideostreampool-tags
	//
	Tags *[]*CfnMediaPipelineKinesisVideoStreamPoolPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}


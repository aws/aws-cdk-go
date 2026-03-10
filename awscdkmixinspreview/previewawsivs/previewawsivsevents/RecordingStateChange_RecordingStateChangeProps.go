package previewawsivsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.ivs@RecordingStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recordingStateChangeProps := &RecordingStateChangeProps{
//   	ChannelName: []*string{
//   		jsii.String("channelName"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	RecordingS3BucketName: []*string{
//   		jsii.String("recordingS3BucketName"),
//   	},
//   	RecordingS3KeyPrefix: []*string{
//   		jsii.String("recordingS3KeyPrefix"),
//   	},
//   	RecordingStatus: []*string{
//   		jsii.String("recordingStatus"),
//   	},
//   	RecordingStatusReason: []*string{
//   		jsii.String("recordingStatusReason"),
//   	},
//   	StreamId: []*string{
//   		jsii.String("streamId"),
//   	},
//   }
//
// Experimental.
type RecordingStateChange_RecordingStateChangeProps struct {
	// channel_name property.
	//
	// Specify an array of string values to match this event if the actual value of channel_name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChannelName *[]*string `field:"optional" json:"channelName" yaml:"channelName"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// recording_s3_bucket_name property.
	//
	// Specify an array of string values to match this event if the actual value of recording_s3_bucket_name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RecordingS3BucketName *[]*string `field:"optional" json:"recordingS3BucketName" yaml:"recordingS3BucketName"`
	// recording_s3_key_prefix property.
	//
	// Specify an array of string values to match this event if the actual value of recording_s3_key_prefix is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RecordingS3KeyPrefix *[]*string `field:"optional" json:"recordingS3KeyPrefix" yaml:"recordingS3KeyPrefix"`
	// recording_status property.
	//
	// Specify an array of string values to match this event if the actual value of recording_status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RecordingStatus *[]*string `field:"optional" json:"recordingStatus" yaml:"recordingStatus"`
	// recording_status_reason property.
	//
	// Specify an array of string values to match this event if the actual value of recording_status_reason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RecordingStatusReason *[]*string `field:"optional" json:"recordingStatusReason" yaml:"recordingStatusReason"`
	// stream_id property.
	//
	// Specify an array of string values to match this event if the actual value of stream_id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StreamId *[]*string `field:"optional" json:"streamId" yaml:"streamId"`
}


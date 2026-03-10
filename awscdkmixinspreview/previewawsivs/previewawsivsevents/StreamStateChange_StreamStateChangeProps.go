package previewawsivsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.ivs@StreamStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   streamStateChangeProps := &StreamStateChangeProps{
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
//   	EventName: []*string{
//   		jsii.String("eventName"),
//   	},
//   	StreamId: []*string{
//   		jsii.String("streamId"),
//   	},
//   }
//
// Experimental.
type StreamStateChange_StreamStateChangeProps struct {
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
	// event_name property.
	//
	// Specify an array of string values to match this event if the actual value of event_name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventName *[]*string `field:"optional" json:"eventName" yaml:"eventName"`
	// stream_id property.
	//
	// Specify an array of string values to match this event if the actual value of stream_id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StreamId *[]*string `field:"optional" json:"streamId" yaml:"streamId"`
}


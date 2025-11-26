package previewawsmedialiveevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Channel aws.medialive@MediaLiveChannelInputChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaLiveChannelInputChangeProps := &MediaLiveChannelInputChangeProps{
//   	ActiveInputAttachmentName: []*string{
//   		jsii.String("activeInputAttachmentName"),
//   	},
//   	ActiveInputSwitchActionName: []*string{
//   		jsii.String("activeInputSwitchActionName"),
//   	},
//   	ChannelArn: []*string{
//   		jsii.String("channelArn"),
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
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   	Pipeline: []*string{
//   		jsii.String("pipeline"),
//   	},
//   }
//
// Experimental.
type ChannelEvents_MediaLiveChannelInputChange_MediaLiveChannelInputChangeProps struct {
	// active_input_attachment_name property.
	//
	// Specify an array of string values to match this event if the actual value of active_input_attachment_name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActiveInputAttachmentName *[]*string `field:"optional" json:"activeInputAttachmentName" yaml:"activeInputAttachmentName"`
	// active_input_switch_action_name property.
	//
	// Specify an array of string values to match this event if the actual value of active_input_switch_action_name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActiveInputSwitchActionName *[]*string `field:"optional" json:"activeInputSwitchActionName" yaml:"activeInputSwitchActionName"`
	// channel_arn property.
	//
	// Specify an array of string values to match this event if the actual value of channel_arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Channel reference.
	//
	// Experimental.
	ChannelArn *[]*string `field:"optional" json:"channelArn" yaml:"channelArn"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// message property.
	//
	// Specify an array of string values to match this event if the actual value of message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
	// pipeline property.
	//
	// Specify an array of string values to match this event if the actual value of pipeline is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Pipeline *[]*string `field:"optional" json:"pipeline" yaml:"pipeline"`
}


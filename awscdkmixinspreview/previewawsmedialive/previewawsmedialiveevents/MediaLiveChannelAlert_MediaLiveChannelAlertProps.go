package previewawsmedialiveevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.medialive@MediaLiveChannelAlert event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaLiveChannelAlertProps := &MediaLiveChannelAlertProps{
//   	AlarmId: []*string{
//   		jsii.String("alarmId"),
//   	},
//   	AlarmState: []*string{
//   		jsii.String("alarmState"),
//   	},
//   	AlertType: []*string{
//   		jsii.String("alertType"),
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
type MediaLiveChannelAlert_MediaLiveChannelAlertProps struct {
	// alarm_id property.
	//
	// Specify an array of string values to match this event if the actual value of alarm_id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AlarmId *[]*string `field:"optional" json:"alarmId" yaml:"alarmId"`
	// alarm_state property.
	//
	// Specify an array of string values to match this event if the actual value of alarm_state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AlarmState *[]*string `field:"optional" json:"alarmState" yaml:"alarmState"`
	// alert_type property.
	//
	// Specify an array of string values to match this event if the actual value of alert_type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AlertType *[]*string `field:"optional" json:"alertType" yaml:"alertType"`
	// channel_arn property.
	//
	// Specify an array of string values to match this event if the actual value of channel_arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
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


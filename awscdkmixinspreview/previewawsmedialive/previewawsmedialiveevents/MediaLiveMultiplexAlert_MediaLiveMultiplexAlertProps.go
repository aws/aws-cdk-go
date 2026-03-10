package previewawsmedialiveevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.medialive@MediaLiveMultiplexAlert event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaLiveMultiplexAlertProps := &MediaLiveMultiplexAlertProps{
//   	AlarmState: []*string{
//   		jsii.String("alarmState"),
//   	},
//   	AlertType: []*string{
//   		jsii.String("alertType"),
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
type MediaLiveMultiplexAlert_MediaLiveMultiplexAlertProps struct {
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


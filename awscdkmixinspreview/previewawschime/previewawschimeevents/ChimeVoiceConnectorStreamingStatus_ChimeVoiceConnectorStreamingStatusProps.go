package previewawschimeevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.chime@ChimeVoiceConnectorStreamingStatus event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   chimeVoiceConnectorStreamingStatusProps := &ChimeVoiceConnectorStreamingStatusProps{
//   	CallId: []*string{
//   		jsii.String("callId"),
//   	},
//   	Direction: []*string{
//   		jsii.String("direction"),
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
//   	InviteHeaders: map[string]*string{
//   		"inviteHeadersKey": jsii.String("inviteHeaders"),
//   	},
//   	MediaType: []*string{
//   		jsii.String("mediaType"),
//   	},
//   	SiprecMetadata: []*string{
//   		jsii.String("siprecMetadata"),
//   	},
//   	StartFragmentNumber: []*string{
//   		jsii.String("startFragmentNumber"),
//   	},
//   	StartTime: []*string{
//   		jsii.String("startTime"),
//   	},
//   	StreamArn: []*string{
//   		jsii.String("streamArn"),
//   	},
//   	StreamingStatus: []*string{
//   		jsii.String("streamingStatus"),
//   	},
//   	TransactionId: []*string{
//   		jsii.String("transactionId"),
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   	VoiceConnectorId: []*string{
//   		jsii.String("voiceConnectorId"),
//   	},
//   }
//
// Experimental.
type ChimeVoiceConnectorStreamingStatus_ChimeVoiceConnectorStreamingStatusProps struct {
	// callId property.
	//
	// Specify an array of string values to match this event if the actual value of callId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CallId *[]*string `field:"optional" json:"callId" yaml:"callId"`
	// direction property.
	//
	// Specify an array of string values to match this event if the actual value of direction is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Direction *[]*string `field:"optional" json:"direction" yaml:"direction"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// inviteHeaders property.
	//
	// Specify an array of string values to match this event if the actual value of inviteHeaders is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InviteHeaders *map[string]*string `field:"optional" json:"inviteHeaders" yaml:"inviteHeaders"`
	// mediaType property.
	//
	// Specify an array of string values to match this event if the actual value of mediaType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MediaType *[]*string `field:"optional" json:"mediaType" yaml:"mediaType"`
	// siprecMetadata property.
	//
	// Specify an array of string values to match this event if the actual value of siprecMetadata is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SiprecMetadata *[]*string `field:"optional" json:"siprecMetadata" yaml:"siprecMetadata"`
	// startFragmentNumber property.
	//
	// Specify an array of string values to match this event if the actual value of startFragmentNumber is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartFragmentNumber *[]*string `field:"optional" json:"startFragmentNumber" yaml:"startFragmentNumber"`
	// startTime property.
	//
	// Specify an array of string values to match this event if the actual value of startTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTime *[]*string `field:"optional" json:"startTime" yaml:"startTime"`
	// streamArn property.
	//
	// Specify an array of string values to match this event if the actual value of streamArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StreamArn *[]*string `field:"optional" json:"streamArn" yaml:"streamArn"`
	// streamingStatus property.
	//
	// Specify an array of string values to match this event if the actual value of streamingStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StreamingStatus *[]*string `field:"optional" json:"streamingStatus" yaml:"streamingStatus"`
	// transactionId property.
	//
	// Specify an array of string values to match this event if the actual value of transactionId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransactionId *[]*string `field:"optional" json:"transactionId" yaml:"transactionId"`
	// version property.
	//
	// Specify an array of string values to match this event if the actual value of version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
	// voiceConnectorId property.
	//
	// Specify an array of string values to match this event if the actual value of voiceConnectorId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VoiceConnectorId *[]*string `field:"optional" json:"voiceConnectorId" yaml:"voiceConnectorId"`
}


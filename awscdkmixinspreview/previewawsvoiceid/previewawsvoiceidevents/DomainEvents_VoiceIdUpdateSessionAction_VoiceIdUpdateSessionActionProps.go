package previewawsvoiceidevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Domain aws.voiceid@VoiceIdUpdateSessionAction event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdUpdateSessionActionProps := &VoiceIdUpdateSessionActionProps{
//   	Action: []*string{
//   		jsii.String("action"),
//   	},
//   	DomainId: []*string{
//   		jsii.String("domainId"),
//   	},
//   	ErrorInfo: &ErrorInfo{
//   		ErrorCode: []*string{
//   			jsii.String("errorCode"),
//   		},
//   		ErrorMessage: []*string{
//   			jsii.String("errorMessage"),
//   		},
//   		ErrorType: []*string{
//   			jsii.String("errorType"),
//   		},
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
//   	Session: &Session{
//   		AuthenticationConfiguration: &AuthenticationConfiguration{
//   			AcceptanceThreshold: []*string{
//   				jsii.String("acceptanceThreshold"),
//   			},
//   		},
//   		FraudDetectionConfiguration: &FraudDetectionConfiguration{
//   			RiskThreshold: []*string{
//   				jsii.String("riskThreshold"),
//   			},
//   			WatchlistId: []*string{
//   				jsii.String("watchlistId"),
//   			},
//   		},
//   		GeneratedSpeakerId: []*string{
//   			jsii.String("generatedSpeakerId"),
//   		},
//   		SessionId: []*string{
//   			jsii.String("sessionId"),
//   		},
//   		SessionName: []*string{
//   			jsii.String("sessionName"),
//   		},
//   	},
//   	SourceId: []*string{
//   		jsii.String("sourceId"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdUpdateSessionAction_VoiceIdUpdateSessionActionProps struct {
	// action property.
	//
	// Specify an array of string values to match this event if the actual value of action is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Action *[]*string `field:"optional" json:"action" yaml:"action"`
	// domainId property.
	//
	// Specify an array of string values to match this event if the actual value of domainId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Domain reference.
	//
	// Experimental.
	DomainId *[]*string `field:"optional" json:"domainId" yaml:"domainId"`
	// errorInfo property.
	//
	// Specify an array of string values to match this event if the actual value of errorInfo is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorInfo *DomainEvents_VoiceIdUpdateSessionAction_ErrorInfo `field:"optional" json:"errorInfo" yaml:"errorInfo"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// session property.
	//
	// Specify an array of string values to match this event if the actual value of session is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Session *DomainEvents_VoiceIdUpdateSessionAction_Session `field:"optional" json:"session" yaml:"session"`
	// sourceId property.
	//
	// Specify an array of string values to match this event if the actual value of sourceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceId *[]*string `field:"optional" json:"sourceId" yaml:"sourceId"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
}


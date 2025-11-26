package previewawsvoiceidevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Domain aws.voiceid@VoiceIdFraudsterAction event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdFraudsterActionProps := &VoiceIdFraudsterActionProps{
//   	Action: []*string{
//   		jsii.String("action"),
//   	},
//   	Data: &Data{
//   		RegistrationSource: []*string{
//   			jsii.String("registrationSource"),
//   		},
//   		RegistrationSourceId: []*string{
//   			jsii.String("registrationSourceId"),
//   		},
//   		RegistrationStatus: []*string{
//   			jsii.String("registrationStatus"),
//   		},
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
//   	GeneratedFraudsterId: []*string{
//   		jsii.String("generatedFraudsterId"),
//   	},
//   	SourceId: []*string{
//   		jsii.String("sourceId"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   	WatchlistIds: []*string{
//   		jsii.String("watchlistIds"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdFraudsterAction_VoiceIdFraudsterActionProps struct {
	// action property.
	//
	// Specify an array of string values to match this event if the actual value of action is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Action *[]*string `field:"optional" json:"action" yaml:"action"`
	// data property.
	//
	// Specify an array of string values to match this event if the actual value of data is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Data *DomainEvents_VoiceIdFraudsterAction_Data `field:"optional" json:"data" yaml:"data"`
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
	ErrorInfo *DomainEvents_VoiceIdFraudsterAction_ErrorInfo `field:"optional" json:"errorInfo" yaml:"errorInfo"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// generatedFraudsterId property.
	//
	// Specify an array of string values to match this event if the actual value of generatedFraudsterId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GeneratedFraudsterId *[]*string `field:"optional" json:"generatedFraudsterId" yaml:"generatedFraudsterId"`
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
	// watchlistIds property.
	//
	// Specify an array of string values to match this event if the actual value of watchlistIds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	WatchlistIds *[]*string `field:"optional" json:"watchlistIds" yaml:"watchlistIds"`
}


package previewawsvoiceidevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Domain aws.voiceid@VoiceIdBatchFraudsterRegistrationAction event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceIdBatchFraudsterRegistrationActionProps := &VoiceIdBatchFraudsterRegistrationActionProps{
//   	Action: []*string{
//   		jsii.String("action"),
//   	},
//   	BatchJobId: []*string{
//   		jsii.String("batchJobId"),
//   	},
//   	Data: &Data{
//   		DataAccessRoleArn: []*string{
//   			jsii.String("dataAccessRoleArn"),
//   		},
//   		InputDataConfig: &InputDataConfig{
//   			S3Uri: []*string{
//   				jsii.String("s3Uri"),
//   			},
//   		},
//   		OutputDataConfig: &OutputDataConfig{
//   			KmsKeyId: []*string{
//   				jsii.String("kmsKeyId"),
//   			},
//   			S3Uri: []*string{
//   				jsii.String("s3Uri"),
//   			},
//   		},
//   		RegistrationConfig: &RegistrationConfig{
//   			DuplicateRegistrationAction: []*string{
//   				jsii.String("duplicateRegistrationAction"),
//   			},
//   			FraudsterSimilarityThreshold: []*string{
//   				jsii.String("fraudsterSimilarityThreshold"),
//   			},
//   			WatchlistIds: []*string{
//   				jsii.String("watchlistIds"),
//   			},
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
//   	SourceId: []*string{
//   		jsii.String("sourceId"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdBatchFraudsterRegistrationAction_VoiceIdBatchFraudsterRegistrationActionProps struct {
	// action property.
	//
	// Specify an array of string values to match this event if the actual value of action is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Action *[]*string `field:"optional" json:"action" yaml:"action"`
	// batchJobId property.
	//
	// Specify an array of string values to match this event if the actual value of batchJobId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BatchJobId *[]*string `field:"optional" json:"batchJobId" yaml:"batchJobId"`
	// data property.
	//
	// Specify an array of string values to match this event if the actual value of data is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Data *DomainEvents_VoiceIdBatchFraudsterRegistrationAction_Data `field:"optional" json:"data" yaml:"data"`
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
	ErrorInfo *DomainEvents_VoiceIdBatchFraudsterRegistrationAction_ErrorInfo `field:"optional" json:"errorInfo" yaml:"errorInfo"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
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


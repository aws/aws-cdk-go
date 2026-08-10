package previewawstranscribeevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.transcribe@TranscribeJobStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transcribeJobStateChangeProps := &TranscribeJobStateChangeProps{
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
//   	TranscriptionJobName: []*string{
//   		jsii.String("transcriptionJobName"),
//   	},
//   	TranscriptionJobStatus: []*string{
//   		jsii.String("transcriptionJobStatus"),
//   	},
//   }
//
// Experimental.
type TranscribeJobStateChange_TranscribeJobStateChangeProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// TranscriptionJobName property.
	//
	// Specify an array of string values to match this event if the actual value of TranscriptionJobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TranscriptionJobName *[]*string `field:"optional" json:"transcriptionJobName" yaml:"transcriptionJobName"`
	// TranscriptionJobStatus property.
	//
	// Specify an array of string values to match this event if the actual value of TranscriptionJobStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TranscriptionJobStatus *[]*string `field:"optional" json:"transcriptionJobStatus" yaml:"transcriptionJobStatus"`
}


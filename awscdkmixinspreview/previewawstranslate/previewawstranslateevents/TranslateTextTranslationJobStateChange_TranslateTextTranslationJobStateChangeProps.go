package previewawstranslateevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.translate@TranslateTextTranslationJobStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   translateTextTranslationJobStateChangeProps := &TranslateTextTranslationJobStateChangeProps{
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
//   	LatestUpdateAttemptAt: []*string{
//   		jsii.String("latestUpdateAttemptAt"),
//   	},
//   	LatestUpdateAttemptStatus: []*string{
//   		jsii.String("latestUpdateAttemptStatus"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Operation: []*string{
//   		jsii.String("operation"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   }
//
// Experimental.
type TranslateTextTranslationJobStateChange_TranslateTextTranslationJobStateChangeProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// latestUpdateAttemptAt property.
	//
	// Specify an array of string values to match this event if the actual value of latestUpdateAttemptAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LatestUpdateAttemptAt *[]*string `field:"optional" json:"latestUpdateAttemptAt" yaml:"latestUpdateAttemptAt"`
	// latestUpdateAttemptStatus property.
	//
	// Specify an array of string values to match this event if the actual value of latestUpdateAttemptStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LatestUpdateAttemptStatus *[]*string `field:"optional" json:"latestUpdateAttemptStatus" yaml:"latestUpdateAttemptStatus"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// operation property.
	//
	// Specify an array of string values to match this event if the actual value of operation is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Operation *[]*string `field:"optional" json:"operation" yaml:"operation"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
}


package previewawshealthlakeevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.healthlake@ExportJobCompleted event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   exportJobCompletedProps := &ExportJobCompletedProps{
//   	DatastoreId: []*string{
//   		jsii.String("datastoreId"),
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
//   	JobId: []*string{
//   		jsii.String("jobId"),
//   	},
//   	OutputDataConfig: &OutputDataConfig{
//   		S3Uri: []*string{
//   			jsii.String("s3Uri"),
//   		},
//   	},
//   	SubmitTime: []*string{
//   		jsii.String("submitTime"),
//   	},
//   }
//
// Experimental.
type ExportJobCompleted_ExportJobCompletedProps struct {
	// datastoreId property.
	//
	// Specify an array of string values to match this event if the actual value of datastoreId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatastoreId *[]*string `field:"optional" json:"datastoreId" yaml:"datastoreId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// jobId property.
	//
	// Specify an array of string values to match this event if the actual value of jobId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobId *[]*string `field:"optional" json:"jobId" yaml:"jobId"`
	// outputDataConfig property.
	//
	// Specify an array of string values to match this event if the actual value of outputDataConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OutputDataConfig *ExportJobCompleted_OutputDataConfig `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// submitTime property.
	//
	// Specify an array of string values to match this event if the actual value of submitTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SubmitTime *[]*string `field:"optional" json:"submitTime" yaml:"submitTime"`
}


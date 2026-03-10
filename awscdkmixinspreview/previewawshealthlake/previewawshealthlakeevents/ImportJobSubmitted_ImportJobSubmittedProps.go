package previewawshealthlakeevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.healthlake@ImportJobSubmitted event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   importJobSubmittedProps := &ImportJobSubmittedProps{
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
//   	InputDataConfig: &InputDataConfig{
//   		S3Uri: []*string{
//   			jsii.String("s3Uri"),
//   		},
//   	},
//   	JobId: []*string{
//   		jsii.String("jobId"),
//   	},
//   	SubmitTime: []*string{
//   		jsii.String("submitTime"),
//   	},
//   }
//
// Experimental.
type ImportJobSubmitted_ImportJobSubmittedProps struct {
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
	// inputDataConfig property.
	//
	// Specify an array of string values to match this event if the actual value of inputDataConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InputDataConfig *ImportJobSubmitted_InputDataConfig `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// jobId property.
	//
	// Specify an array of string values to match this event if the actual value of jobId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobId *[]*string `field:"optional" json:"jobId" yaml:"jobId"`
	// submitTime property.
	//
	// Specify an array of string values to match this event if the actual value of submitTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SubmitTime *[]*string `field:"optional" json:"submitTime" yaml:"submitTime"`
}


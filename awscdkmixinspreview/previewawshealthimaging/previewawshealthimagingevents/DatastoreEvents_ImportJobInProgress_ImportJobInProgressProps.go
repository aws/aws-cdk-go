package previewawshealthimagingevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Datastore aws.healthimaging@ImportJobInProgress event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   importJobInProgressProps := &ImportJobInProgressProps{
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
//   	ImagingVersion: []*string{
//   		jsii.String("imagingVersion"),
//   	},
//   	InputS3Uri: []*string{
//   		jsii.String("inputS3Uri"),
//   	},
//   	JobId: []*string{
//   		jsii.String("jobId"),
//   	},
//   	JobName: []*string{
//   		jsii.String("jobName"),
//   	},
//   	JobStatus: []*string{
//   		jsii.String("jobStatus"),
//   	},
//   	OutputS3Uri: []*string{
//   		jsii.String("outputS3Uri"),
//   	},
//   }
//
// Experimental.
type DatastoreEvents_ImportJobInProgress_ImportJobInProgressProps struct {
	// datastoreId property.
	//
	// Specify an array of string values to match this event if the actual value of datastoreId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Datastore reference.
	//
	// Experimental.
	DatastoreId *[]*string `field:"optional" json:"datastoreId" yaml:"datastoreId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// imagingVersion property.
	//
	// Specify an array of string values to match this event if the actual value of imagingVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImagingVersion *[]*string `field:"optional" json:"imagingVersion" yaml:"imagingVersion"`
	// inputS3Uri property.
	//
	// Specify an array of string values to match this event if the actual value of inputS3Uri is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InputS3Uri *[]*string `field:"optional" json:"inputS3Uri" yaml:"inputS3Uri"`
	// jobId property.
	//
	// Specify an array of string values to match this event if the actual value of jobId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobId *[]*string `field:"optional" json:"jobId" yaml:"jobId"`
	// jobName property.
	//
	// Specify an array of string values to match this event if the actual value of jobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobName *[]*string `field:"optional" json:"jobName" yaml:"jobName"`
	// jobStatus property.
	//
	// Specify an array of string values to match this event if the actual value of jobStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobStatus *[]*string `field:"optional" json:"jobStatus" yaml:"jobStatus"`
	// outputS3Uri property.
	//
	// Specify an array of string values to match this event if the actual value of outputS3Uri is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OutputS3Uri *[]*string `field:"optional" json:"outputS3Uri" yaml:"outputS3Uri"`
}


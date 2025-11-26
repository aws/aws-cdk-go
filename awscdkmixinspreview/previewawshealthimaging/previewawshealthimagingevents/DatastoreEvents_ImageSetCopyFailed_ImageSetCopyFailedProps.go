package previewawshealthimagingevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Datastore aws.healthimaging@ImageSetCopyFailed event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetCopyFailedProps := &ImageSetCopyFailedProps{
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
//   	ImageSetId: []*string{
//   		jsii.String("imageSetId"),
//   	},
//   	ImageSetState: []*string{
//   		jsii.String("imageSetState"),
//   	},
//   	ImageSetWorkflowStatus: []*string{
//   		jsii.String("imageSetWorkflowStatus"),
//   	},
//   	ImagingVersion: []*string{
//   		jsii.String("imagingVersion"),
//   	},
//   }
//
// Experimental.
type DatastoreEvents_ImageSetCopyFailed_ImageSetCopyFailedProps struct {
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
	// imageSetId property.
	//
	// Specify an array of string values to match this event if the actual value of imageSetId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageSetId *[]*string `field:"optional" json:"imageSetId" yaml:"imageSetId"`
	// imageSetState property.
	//
	// Specify an array of string values to match this event if the actual value of imageSetState is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageSetState *[]*string `field:"optional" json:"imageSetState" yaml:"imageSetState"`
	// imageSetWorkflowStatus property.
	//
	// Specify an array of string values to match this event if the actual value of imageSetWorkflowStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageSetWorkflowStatus *[]*string `field:"optional" json:"imageSetWorkflowStatus" yaml:"imageSetWorkflowStatus"`
	// imagingVersion property.
	//
	// Specify an array of string values to match this event if the actual value of imagingVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImagingVersion *[]*string `field:"optional" json:"imagingVersion" yaml:"imagingVersion"`
}


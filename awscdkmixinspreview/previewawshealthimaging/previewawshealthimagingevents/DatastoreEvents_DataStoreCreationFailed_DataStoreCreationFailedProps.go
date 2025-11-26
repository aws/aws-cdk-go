package previewawshealthimagingevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Datastore aws.healthimaging@DataStoreCreationFailed event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStoreCreationFailedProps := &DataStoreCreationFailedProps{
//   	DatastoreId: []*string{
//   		jsii.String("datastoreId"),
//   	},
//   	DatastoreName: []*string{
//   		jsii.String("datastoreName"),
//   	},
//   	DatastoreStatus: []*string{
//   		jsii.String("datastoreStatus"),
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
//   }
//
// Experimental.
type DatastoreEvents_DataStoreCreationFailed_DataStoreCreationFailedProps struct {
	// datastoreId property.
	//
	// Specify an array of string values to match this event if the actual value of datastoreId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Datastore reference.
	//
	// Experimental.
	DatastoreId *[]*string `field:"optional" json:"datastoreId" yaml:"datastoreId"`
	// datastoreName property.
	//
	// Specify an array of string values to match this event if the actual value of datastoreName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatastoreName *[]*string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// datastoreStatus property.
	//
	// Specify an array of string values to match this event if the actual value of datastoreStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatastoreStatus *[]*string `field:"optional" json:"datastoreStatus" yaml:"datastoreStatus"`
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
}


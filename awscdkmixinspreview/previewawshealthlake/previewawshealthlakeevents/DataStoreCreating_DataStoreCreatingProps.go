package previewawshealthlakeevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.healthlake@DataStoreCreating event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStoreCreatingProps := &DataStoreCreatingProps{
//   	DatastoreEndpoint: []*string{
//   		jsii.String("datastoreEndpoint"),
//   	},
//   	DatastoreId: []*string{
//   		jsii.String("datastoreId"),
//   	},
//   	DatastoreName: []*string{
//   		jsii.String("datastoreName"),
//   	},
//   	DatastoreTypeVersion: []*string{
//   		jsii.String("datastoreTypeVersion"),
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
//   }
//
// Experimental.
type DataStoreCreating_DataStoreCreatingProps struct {
	// datastoreEndpoint property.
	//
	// Specify an array of string values to match this event if the actual value of datastoreEndpoint is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatastoreEndpoint *[]*string `field:"optional" json:"datastoreEndpoint" yaml:"datastoreEndpoint"`
	// datastoreId property.
	//
	// Specify an array of string values to match this event if the actual value of datastoreId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
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
	// datastoreTypeVersion property.
	//
	// Specify an array of string values to match this event if the actual value of datastoreTypeVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatastoreTypeVersion *[]*string `field:"optional" json:"datastoreTypeVersion" yaml:"datastoreTypeVersion"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
}


package previewawsstoragegatewayevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.storagegateway@StorageGatewayObjectUploadEvent event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   storageGatewayObjectUploadEventProps := &StorageGatewayObjectUploadEventProps{
//   	BucketName: []*string{
//   		jsii.String("bucketName"),
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
//   	EventType: []*string{
//   		jsii.String("eventType"),
//   	},
//   	ModificationTime: []*string{
//   		jsii.String("modificationTime"),
//   	},
//   	ObjectKey: []*string{
//   		jsii.String("objectKey"),
//   	},
//   	ObjectSize: []*string{
//   		jsii.String("objectSize"),
//   	},
//   	Prefix: []*string{
//   		jsii.String("prefix"),
//   	},
//   }
//
// Experimental.
type StorageGatewayObjectUploadEvent_StorageGatewayObjectUploadEventProps struct {
	// bucket-name property.
	//
	// Specify an array of string values to match this event if the actual value of bucket-name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BucketName *[]*string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// event-type property.
	//
	// Specify an array of string values to match this event if the actual value of event-type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// modification-time property.
	//
	// Specify an array of string values to match this event if the actual value of modification-time is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ModificationTime *[]*string `field:"optional" json:"modificationTime" yaml:"modificationTime"`
	// object-key property.
	//
	// Specify an array of string values to match this event if the actual value of object-key is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ObjectKey *[]*string `field:"optional" json:"objectKey" yaml:"objectKey"`
	// object-size property.
	//
	// Specify an array of string values to match this event if the actual value of object-size is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ObjectSize *[]*string `field:"optional" json:"objectSize" yaml:"objectSize"`
	// prefix property.
	//
	// Specify an array of string values to match this event if the actual value of prefix is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Prefix *[]*string `field:"optional" json:"prefix" yaml:"prefix"`
}


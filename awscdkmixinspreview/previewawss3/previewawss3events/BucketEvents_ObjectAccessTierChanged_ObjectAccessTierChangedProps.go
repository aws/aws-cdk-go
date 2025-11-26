package previewawss3events

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Bucket aws.s3@ObjectAccessTierChanged event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectAccessTierChangedProps := &ObjectAccessTierChangedProps{
//   	Bucket: &Bucket{
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   	},
//   	DestinationAccessTier: []*string{
//   		jsii.String("destinationAccessTier"),
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
//   	Object: &ObjectType{
//   		Etag: []*string{
//   			jsii.String("etag"),
//   		},
//   		Key: []*string{
//   			jsii.String("key"),
//   		},
//   		Size: []*string{
//   			jsii.String("size"),
//   		},
//   		VersionId: []*string{
//   			jsii.String("versionId"),
//   		},
//   	},
//   	Requester: []*string{
//   		jsii.String("requester"),
//   	},
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   }
//
// Experimental.
type BucketEvents_ObjectAccessTierChanged_ObjectAccessTierChangedProps struct {
	// bucket property.
	//
	// Specify an array of string values to match this event if the actual value of bucket is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Bucket *BucketEvents_ObjectAccessTierChanged_Bucket `field:"optional" json:"bucket" yaml:"bucket"`
	// destination-access-tier property.
	//
	// Specify an array of string values to match this event if the actual value of destination-access-tier is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DestinationAccessTier *[]*string `field:"optional" json:"destinationAccessTier" yaml:"destinationAccessTier"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// object property.
	//
	// Specify an array of string values to match this event if the actual value of object is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Object *BucketEvents_ObjectAccessTierChanged_ObjectType `field:"optional" json:"object" yaml:"object"`
	// requester property.
	//
	// Specify an array of string values to match this event if the actual value of requester is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Requester *[]*string `field:"optional" json:"requester" yaml:"requester"`
	// request-id property.
	//
	// Specify an array of string values to match this event if the actual value of request-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// version property.
	//
	// Specify an array of string values to match this event if the actual value of version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}


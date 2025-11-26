package previewawss3events

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Bucket aws.s3@ObjectDeleted event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectDeletedProps := &ObjectDeletedProps{
//   	Bucket: &Bucket{
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   	},
//   	DeletionType: []*string{
//   		jsii.String("deletionType"),
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
//   		Sequencer: []*string{
//   			jsii.String("sequencer"),
//   		},
//   		VersionId: []*string{
//   			jsii.String("versionId"),
//   		},
//   	},
//   	Reason: []*string{
//   		jsii.String("reason"),
//   	},
//   	Requester: []*string{
//   		jsii.String("requester"),
//   	},
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	SourceIpAddress: []*string{
//   		jsii.String("sourceIpAddress"),
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   }
//
// Experimental.
type BucketEvents_ObjectDeleted_ObjectDeletedProps struct {
	// bucket property.
	//
	// Specify an array of string values to match this event if the actual value of bucket is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Bucket *BucketEvents_ObjectDeleted_Bucket `field:"optional" json:"bucket" yaml:"bucket"`
	// deletion-type property.
	//
	// Specify an array of string values to match this event if the actual value of deletion-type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeletionType *[]*string `field:"optional" json:"deletionType" yaml:"deletionType"`
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
	Object *BucketEvents_ObjectDeleted_ObjectType `field:"optional" json:"object" yaml:"object"`
	// reason property.
	//
	// Specify an array of string values to match this event if the actual value of reason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Reason *[]*string `field:"optional" json:"reason" yaml:"reason"`
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
	// source-ip-address property.
	//
	// Specify an array of string values to match this event if the actual value of source-ip-address is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceIpAddress *[]*string `field:"optional" json:"sourceIpAddress" yaml:"sourceIpAddress"`
	// version property.
	//
	// Specify an array of string values to match this event if the actual value of version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}


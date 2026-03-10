package previewawsomicsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.omics@S3AccessPolicyStatusChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3AccessPolicyStatusChangeProps := &S3AccessPolicyStatusChangeProps{
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
//   	OmicsVersion: []*string{
//   		jsii.String("omicsVersion"),
//   	},
//   	S3AccessPointArn: []*string{
//   		jsii.String("s3AccessPointArn"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   	StoreId: []*string{
//   		jsii.String("storeId"),
//   	},
//   	StoreType: []*string{
//   		jsii.String("storeType"),
//   	},
//   }
//
// Experimental.
type S3AccessPolicyStatusChange_S3AccessPolicyStatusChangeProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// omicsVersion property.
	//
	// Specify an array of string values to match this event if the actual value of omicsVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OmicsVersion *[]*string `field:"optional" json:"omicsVersion" yaml:"omicsVersion"`
	// s3AccessPointArn property.
	//
	// Specify an array of string values to match this event if the actual value of s3AccessPointArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3AccessPointArn *[]*string `field:"optional" json:"s3AccessPointArn" yaml:"s3AccessPointArn"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
	// storeId property.
	//
	// Specify an array of string values to match this event if the actual value of storeId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StoreId *[]*string `field:"optional" json:"storeId" yaml:"storeId"`
	// storeType property.
	//
	// Specify an array of string values to match this event if the actual value of storeType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StoreType *[]*string `field:"optional" json:"storeType" yaml:"storeType"`
}


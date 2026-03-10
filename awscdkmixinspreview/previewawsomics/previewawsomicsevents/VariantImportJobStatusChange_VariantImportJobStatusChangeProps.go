package previewawsomicsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.omics@VariantImportJobStatusChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   variantImportJobStatusChangeProps := &VariantImportJobStatusChangeProps{
//   	Arn: []*string{
//   		jsii.String("arn"),
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
//   	OmicsVersion: []*string{
//   		jsii.String("omicsVersion"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   	StoreId: []*string{
//   		jsii.String("storeId"),
//   	},
//   	StoreName: []*string{
//   		jsii.String("storeName"),
//   	},
//   }
//
// Experimental.
type VariantImportJobStatusChange_VariantImportJobStatusChangeProps struct {
	// arn property.
	//
	// Specify an array of string values to match this event if the actual value of arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
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
	// omicsVersion property.
	//
	// Specify an array of string values to match this event if the actual value of omicsVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OmicsVersion *[]*string `field:"optional" json:"omicsVersion" yaml:"omicsVersion"`
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
	// storeName property.
	//
	// Specify an array of string values to match this event if the actual value of storeName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StoreName *[]*string `field:"optional" json:"storeName" yaml:"storeName"`
}


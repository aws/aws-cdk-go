package previewawsomicsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for ReferenceStore aws.omics@ReferenceStatusChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceStatusChangeProps := &ReferenceStatusChangeProps{
//   	Arn: []*string{
//   		jsii.String("arn"),
//   	},
//   	CreationJobId: []*string{
//   		jsii.String("creationJobId"),
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
//   	Id: []*string{
//   		jsii.String("id"),
//   	},
//   	OmicsVersion: []*string{
//   		jsii.String("omicsVersion"),
//   	},
//   	ReferenceStoreId: []*string{
//   		jsii.String("referenceStoreId"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   }
//
// Experimental.
type ReferenceStoreEvents_ReferenceStatusChange_ReferenceStatusChangeProps struct {
	// arn property.
	//
	// Specify an array of string values to match this event if the actual value of arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
	// creationJobId property.
	//
	// Specify an array of string values to match this event if the actual value of creationJobId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreationJobId *[]*string `field:"optional" json:"creationJobId" yaml:"creationJobId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// id property.
	//
	// Specify an array of string values to match this event if the actual value of id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Id *[]*string `field:"optional" json:"id" yaml:"id"`
	// omicsVersion property.
	//
	// Specify an array of string values to match this event if the actual value of omicsVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OmicsVersion *[]*string `field:"optional" json:"omicsVersion" yaml:"omicsVersion"`
	// referenceStoreId property.
	//
	// Specify an array of string values to match this event if the actual value of referenceStoreId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the ReferenceStore reference.
	//
	// Experimental.
	ReferenceStoreId *[]*string `field:"optional" json:"referenceStoreId" yaml:"referenceStoreId"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
}


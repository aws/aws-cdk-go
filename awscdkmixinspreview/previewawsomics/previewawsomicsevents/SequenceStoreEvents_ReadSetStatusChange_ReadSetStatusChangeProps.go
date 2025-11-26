package previewawsomicsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for SequenceStore aws.omics@ReadSetStatusChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   readSetStatusChangeProps := &ReadSetStatusChangeProps{
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
//   	SequenceStoreId: []*string{
//   		jsii.String("sequenceStoreId"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   	StatusMessage: []*string{
//   		jsii.String("statusMessage"),
//   	},
//   }
//
// Experimental.
type SequenceStoreEvents_ReadSetStatusChange_ReadSetStatusChangeProps struct {
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
	// sequenceStoreId property.
	//
	// Specify an array of string values to match this event if the actual value of sequenceStoreId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the SequenceStore reference.
	//
	// Experimental.
	SequenceStoreId *[]*string `field:"optional" json:"sequenceStoreId" yaml:"sequenceStoreId"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
	// statusMessage property.
	//
	// Specify an array of string values to match this event if the actual value of statusMessage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StatusMessage *[]*string `field:"optional" json:"statusMessage" yaml:"statusMessage"`
}


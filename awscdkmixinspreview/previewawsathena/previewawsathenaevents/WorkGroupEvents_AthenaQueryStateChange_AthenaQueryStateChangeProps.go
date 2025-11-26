package previewawsathenaevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for WorkGroup aws.athena@AthenaQueryStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   athenaQueryStateChangeProps := &AthenaQueryStateChangeProps{
//   	CurrentState: []*string{
//   		jsii.String("currentState"),
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
//   	PreviousState: []*string{
//   		jsii.String("previousState"),
//   	},
//   	QueryExecutionId: []*string{
//   		jsii.String("queryExecutionId"),
//   	},
//   	SequenceNumber: []*string{
//   		jsii.String("sequenceNumber"),
//   	},
//   	StatementType: []*string{
//   		jsii.String("statementType"),
//   	},
//   	VersionId: []*string{
//   		jsii.String("versionId"),
//   	},
//   	WorkgroupName: []*string{
//   		jsii.String("workgroupName"),
//   	},
//   }
//
// Experimental.
type WorkGroupEvents_AthenaQueryStateChange_AthenaQueryStateChangeProps struct {
	// currentState property.
	//
	// Specify an array of string values to match this event if the actual value of currentState is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CurrentState *[]*string `field:"optional" json:"currentState" yaml:"currentState"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// previousState property.
	//
	// Specify an array of string values to match this event if the actual value of previousState is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousState *[]*string `field:"optional" json:"previousState" yaml:"previousState"`
	// queryExecutionId property.
	//
	// Specify an array of string values to match this event if the actual value of queryExecutionId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	QueryExecutionId *[]*string `field:"optional" json:"queryExecutionId" yaml:"queryExecutionId"`
	// sequenceNumber property.
	//
	// Specify an array of string values to match this event if the actual value of sequenceNumber is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SequenceNumber *[]*string `field:"optional" json:"sequenceNumber" yaml:"sequenceNumber"`
	// statementType property.
	//
	// Specify an array of string values to match this event if the actual value of statementType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StatementType *[]*string `field:"optional" json:"statementType" yaml:"statementType"`
	// versionId property.
	//
	// Specify an array of string values to match this event if the actual value of versionId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VersionId *[]*string `field:"optional" json:"versionId" yaml:"versionId"`
	// workgroupName property.
	//
	// Specify an array of string values to match this event if the actual value of workgroupName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the WorkGroup reference.
	//
	// Experimental.
	WorkgroupName *[]*string `field:"optional" json:"workgroupName" yaml:"workgroupName"`
}


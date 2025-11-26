package previewawscodecommitevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Repository aws.codecommit@CodeCommitRepositoryStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeCommitRepositoryStateChangeProps := &CodeCommitRepositoryStateChangeProps{
//   	BaseCommitId: []*string{
//   		jsii.String("baseCommitId"),
//   	},
//   	CallerUserArn: []*string{
//   		jsii.String("callerUserArn"),
//   	},
//   	CommitId: []*string{
//   		jsii.String("commitId"),
//   	},
//   	ConflictDetailLevel: []*string{
//   		jsii.String("conflictDetailLevel"),
//   	},
//   	ConflictDetailsLevel: []*string{
//   		jsii.String("conflictDetailsLevel"),
//   	},
//   	ConflictResolutionStrategy: []*string{
//   		jsii.String("conflictResolutionStrategy"),
//   	},
//   	DestinationCommitId: []*string{
//   		jsii.String("destinationCommitId"),
//   	},
//   	Event: []*string{
//   		jsii.String("event"),
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
//   	MergeOption: []*string{
//   		jsii.String("mergeOption"),
//   	},
//   	OldCommitId: []*string{
//   		jsii.String("oldCommitId"),
//   	},
//   	ReferenceFullName: []*string{
//   		jsii.String("referenceFullName"),
//   	},
//   	ReferenceName: []*string{
//   		jsii.String("referenceName"),
//   	},
//   	ReferenceType: []*string{
//   		jsii.String("referenceType"),
//   	},
//   	RepositoryId: []*string{
//   		jsii.String("repositoryId"),
//   	},
//   	RepositoryName: []*string{
//   		jsii.String("repositoryName"),
//   	},
//   	SourceCommitId: []*string{
//   		jsii.String("sourceCommitId"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_CodeCommitRepositoryStateChange_CodeCommitRepositoryStateChangeProps struct {
	// baseCommitId property.
	//
	// Specify an array of string values to match this event if the actual value of baseCommitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BaseCommitId *[]*string `field:"optional" json:"baseCommitId" yaml:"baseCommitId"`
	// callerUserArn property.
	//
	// Specify an array of string values to match this event if the actual value of callerUserArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CallerUserArn *[]*string `field:"optional" json:"callerUserArn" yaml:"callerUserArn"`
	// commitId property.
	//
	// Specify an array of string values to match this event if the actual value of commitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CommitId *[]*string `field:"optional" json:"commitId" yaml:"commitId"`
	// conflictDetailLevel property.
	//
	// Specify an array of string values to match this event if the actual value of conflictDetailLevel is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ConflictDetailLevel *[]*string `field:"optional" json:"conflictDetailLevel" yaml:"conflictDetailLevel"`
	// conflictDetailsLevel property.
	//
	// Specify an array of string values to match this event if the actual value of conflictDetailsLevel is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ConflictDetailsLevel *[]*string `field:"optional" json:"conflictDetailsLevel" yaml:"conflictDetailsLevel"`
	// conflictResolutionStrategy property.
	//
	// Specify an array of string values to match this event if the actual value of conflictResolutionStrategy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ConflictResolutionStrategy *[]*string `field:"optional" json:"conflictResolutionStrategy" yaml:"conflictResolutionStrategy"`
	// destinationCommitId property.
	//
	// Specify an array of string values to match this event if the actual value of destinationCommitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DestinationCommitId *[]*string `field:"optional" json:"destinationCommitId" yaml:"destinationCommitId"`
	// event property.
	//
	// Specify an array of string values to match this event if the actual value of event is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Event *[]*string `field:"optional" json:"event" yaml:"event"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// mergeOption property.
	//
	// Specify an array of string values to match this event if the actual value of mergeOption is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MergeOption *[]*string `field:"optional" json:"mergeOption" yaml:"mergeOption"`
	// oldCommitId property.
	//
	// Specify an array of string values to match this event if the actual value of oldCommitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OldCommitId *[]*string `field:"optional" json:"oldCommitId" yaml:"oldCommitId"`
	// referenceFullName property.
	//
	// Specify an array of string values to match this event if the actual value of referenceFullName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReferenceFullName *[]*string `field:"optional" json:"referenceFullName" yaml:"referenceFullName"`
	// referenceName property.
	//
	// Specify an array of string values to match this event if the actual value of referenceName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReferenceName *[]*string `field:"optional" json:"referenceName" yaml:"referenceName"`
	// referenceType property.
	//
	// Specify an array of string values to match this event if the actual value of referenceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReferenceType *[]*string `field:"optional" json:"referenceType" yaml:"referenceType"`
	// repositoryId property.
	//
	// Specify an array of string values to match this event if the actual value of repositoryId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Repository reference.
	//
	// Experimental.
	RepositoryId *[]*string `field:"optional" json:"repositoryId" yaml:"repositoryId"`
	// repositoryName property.
	//
	// Specify an array of string values to match this event if the actual value of repositoryName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RepositoryName *[]*string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// sourceCommitId property.
	//
	// Specify an array of string values to match this event if the actual value of sourceCommitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceCommitId *[]*string `field:"optional" json:"sourceCommitId" yaml:"sourceCommitId"`
}


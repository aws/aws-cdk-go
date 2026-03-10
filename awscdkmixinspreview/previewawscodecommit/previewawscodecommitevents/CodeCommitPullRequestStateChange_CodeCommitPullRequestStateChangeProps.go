package previewawscodecommitevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.codecommit@CodeCommitPullRequestStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeCommitPullRequestStateChangeProps := &CodeCommitPullRequestStateChangeProps{
//   	ApprovalStatus: []*string{
//   		jsii.String("approvalStatus"),
//   	},
//   	Author: []*string{
//   		jsii.String("author"),
//   	},
//   	CallerUserArn: []*string{
//   		jsii.String("callerUserArn"),
//   	},
//   	CreationDate: []*string{
//   		jsii.String("creationDate"),
//   	},
//   	Description: []*string{
//   		jsii.String("description"),
//   	},
//   	DestinationCommit: []*string{
//   		jsii.String("destinationCommit"),
//   	},
//   	DestinationReference: []*string{
//   		jsii.String("destinationReference"),
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
//   	IsMerged: []*string{
//   		jsii.String("isMerged"),
//   	},
//   	LastModifiedDate: []*string{
//   		jsii.String("lastModifiedDate"),
//   	},
//   	MergeOption: []*string{
//   		jsii.String("mergeOption"),
//   	},
//   	NotificationBody: []*string{
//   		jsii.String("notificationBody"),
//   	},
//   	OverrideStatus: []*string{
//   		jsii.String("overrideStatus"),
//   	},
//   	PullRequestId: []*string{
//   		jsii.String("pullRequestId"),
//   	},
//   	PullRequestStatus: []*string{
//   		jsii.String("pullRequestStatus"),
//   	},
//   	RepositoryNames: []*string{
//   		jsii.String("repositoryNames"),
//   	},
//   	RevisionId: []*string{
//   		jsii.String("revisionId"),
//   	},
//   	SourceCommit: []*string{
//   		jsii.String("sourceCommit"),
//   	},
//   	SourceReference: []*string{
//   		jsii.String("sourceReference"),
//   	},
//   	Title: []*string{
//   		jsii.String("title"),
//   	},
//   }
//
// Experimental.
type CodeCommitPullRequestStateChange_CodeCommitPullRequestStateChangeProps struct {
	// approvalStatus property.
	//
	// Specify an array of string values to match this event if the actual value of approvalStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ApprovalStatus *[]*string `field:"optional" json:"approvalStatus" yaml:"approvalStatus"`
	// author property.
	//
	// Specify an array of string values to match this event if the actual value of author is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Author *[]*string `field:"optional" json:"author" yaml:"author"`
	// callerUserArn property.
	//
	// Specify an array of string values to match this event if the actual value of callerUserArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CallerUserArn *[]*string `field:"optional" json:"callerUserArn" yaml:"callerUserArn"`
	// creationDate property.
	//
	// Specify an array of string values to match this event if the actual value of creationDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreationDate *[]*string `field:"optional" json:"creationDate" yaml:"creationDate"`
	// description property.
	//
	// Specify an array of string values to match this event if the actual value of description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// destinationCommit property.
	//
	// Specify an array of string values to match this event if the actual value of destinationCommit is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DestinationCommit *[]*string `field:"optional" json:"destinationCommit" yaml:"destinationCommit"`
	// destinationReference property.
	//
	// Specify an array of string values to match this event if the actual value of destinationReference is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DestinationReference *[]*string `field:"optional" json:"destinationReference" yaml:"destinationReference"`
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
	// isMerged property.
	//
	// Specify an array of string values to match this event if the actual value of isMerged is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IsMerged *[]*string `field:"optional" json:"isMerged" yaml:"isMerged"`
	// lastModifiedDate property.
	//
	// Specify an array of string values to match this event if the actual value of lastModifiedDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastModifiedDate *[]*string `field:"optional" json:"lastModifiedDate" yaml:"lastModifiedDate"`
	// mergeOption property.
	//
	// Specify an array of string values to match this event if the actual value of mergeOption is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MergeOption *[]*string `field:"optional" json:"mergeOption" yaml:"mergeOption"`
	// notificationBody property.
	//
	// Specify an array of string values to match this event if the actual value of notificationBody is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NotificationBody *[]*string `field:"optional" json:"notificationBody" yaml:"notificationBody"`
	// overrideStatus property.
	//
	// Specify an array of string values to match this event if the actual value of overrideStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OverrideStatus *[]*string `field:"optional" json:"overrideStatus" yaml:"overrideStatus"`
	// pullRequestId property.
	//
	// Specify an array of string values to match this event if the actual value of pullRequestId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PullRequestId *[]*string `field:"optional" json:"pullRequestId" yaml:"pullRequestId"`
	// pullRequestStatus property.
	//
	// Specify an array of string values to match this event if the actual value of pullRequestStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PullRequestStatus *[]*string `field:"optional" json:"pullRequestStatus" yaml:"pullRequestStatus"`
	// repositoryNames property.
	//
	// Specify an array of string values to match this event if the actual value of repositoryNames is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RepositoryNames *[]*string `field:"optional" json:"repositoryNames" yaml:"repositoryNames"`
	// revisionId property.
	//
	// Specify an array of string values to match this event if the actual value of revisionId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RevisionId *[]*string `field:"optional" json:"revisionId" yaml:"revisionId"`
	// sourceCommit property.
	//
	// Specify an array of string values to match this event if the actual value of sourceCommit is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceCommit *[]*string `field:"optional" json:"sourceCommit" yaml:"sourceCommit"`
	// sourceReference property.
	//
	// Specify an array of string values to match this event if the actual value of sourceReference is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceReference *[]*string `field:"optional" json:"sourceReference" yaml:"sourceReference"`
	// title property.
	//
	// Specify an array of string values to match this event if the actual value of title is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Title *[]*string `field:"optional" json:"title" yaml:"title"`
}


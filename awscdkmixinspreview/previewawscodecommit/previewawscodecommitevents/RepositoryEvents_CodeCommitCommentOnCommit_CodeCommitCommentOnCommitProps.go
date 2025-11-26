package previewawscodecommitevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Repository aws.codecommit@CodeCommitCommentOnCommit event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeCommitCommentOnCommitProps := &CodeCommitCommentOnCommitProps{
//   	AfterCommitId: []*string{
//   		jsii.String("afterCommitId"),
//   	},
//   	BeforeCommitId: []*string{
//   		jsii.String("beforeCommitId"),
//   	},
//   	CallerUserArn: []*string{
//   		jsii.String("callerUserArn"),
//   	},
//   	CommentId: []*string{
//   		jsii.String("commentId"),
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
//   	InReplyTo: []*string{
//   		jsii.String("inReplyTo"),
//   	},
//   	NotificationBody: []*string{
//   		jsii.String("notificationBody"),
//   	},
//   	RepositoryId: []*string{
//   		jsii.String("repositoryId"),
//   	},
//   	RepositoryName: []*string{
//   		jsii.String("repositoryName"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_CodeCommitCommentOnCommit_CodeCommitCommentOnCommitProps struct {
	// afterCommitId property.
	//
	// Specify an array of string values to match this event if the actual value of afterCommitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AfterCommitId *[]*string `field:"optional" json:"afterCommitId" yaml:"afterCommitId"`
	// beforeCommitId property.
	//
	// Specify an array of string values to match this event if the actual value of beforeCommitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BeforeCommitId *[]*string `field:"optional" json:"beforeCommitId" yaml:"beforeCommitId"`
	// callerUserArn property.
	//
	// Specify an array of string values to match this event if the actual value of callerUserArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CallerUserArn *[]*string `field:"optional" json:"callerUserArn" yaml:"callerUserArn"`
	// commentId property.
	//
	// Specify an array of string values to match this event if the actual value of commentId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CommentId *[]*string `field:"optional" json:"commentId" yaml:"commentId"`
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
	// inReplyTo property.
	//
	// Specify an array of string values to match this event if the actual value of inReplyTo is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InReplyTo *[]*string `field:"optional" json:"inReplyTo" yaml:"inReplyTo"`
	// notificationBody property.
	//
	// Specify an array of string values to match this event if the actual value of notificationBody is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NotificationBody *[]*string `field:"optional" json:"notificationBody" yaml:"notificationBody"`
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
}


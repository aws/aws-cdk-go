package previewawscodecommitevents


// Type definition for PullRequest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var approvalRules interface{}
//
//   pullRequest := &PullRequest{
//   	ApprovalRules: []interface{}{
//   		approvalRules,
//   	},
//   	AuthorArn: []*string{
//   		jsii.String("authorArn"),
//   	},
//   	ClientRequestToken: []*string{
//   		jsii.String("clientRequestToken"),
//   	},
//   	CreationDate: []*string{
//   		jsii.String("creationDate"),
//   	},
//   	Description: []*string{
//   		jsii.String("description"),
//   	},
//   	LastActivityDate: []*string{
//   		jsii.String("lastActivityDate"),
//   	},
//   	PullRequestId: []*string{
//   		jsii.String("pullRequestId"),
//   	},
//   	PullRequestStatus: []*string{
//   		jsii.String("pullRequestStatus"),
//   	},
//   	PullRequestTargets: []PullRequestItem{
//   		&PullRequestItem{
//   			DestinationCommit: []*string{
//   				jsii.String("destinationCommit"),
//   			},
//   			DestinationReference: []*string{
//   				jsii.String("destinationReference"),
//   			},
//   			MergeBase: []*string{
//   				jsii.String("mergeBase"),
//   			},
//   			MergeMetadata: &MergeMetadata{
//   				IsMerged: []*string{
//   					jsii.String("isMerged"),
//   				},
//   				MergeCommitId: []*string{
//   					jsii.String("mergeCommitId"),
//   				},
//   				MergedBy: []*string{
//   					jsii.String("mergedBy"),
//   				},
//   				MergeOption: []*string{
//   					jsii.String("mergeOption"),
//   				},
//   			},
//   			RepositoryName: []*string{
//   				jsii.String("repositoryName"),
//   			},
//   			SourceCommit: []*string{
//   				jsii.String("sourceCommit"),
//   			},
//   			SourceReference: []*string{
//   				jsii.String("sourceReference"),
//   			},
//   		},
//   	},
//   	RevisionId: []*string{
//   		jsii.String("revisionId"),
//   	},
//   	Title: []*string{
//   		jsii.String("title"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_PullRequest struct {
	// approvalRules property.
	//
	// Specify an array of string values to match this event if the actual value of approvalRules is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ApprovalRules *[]interface{} `field:"optional" json:"approvalRules" yaml:"approvalRules"`
	// authorArn property.
	//
	// Specify an array of string values to match this event if the actual value of authorArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AuthorArn *[]*string `field:"optional" json:"authorArn" yaml:"authorArn"`
	// clientRequestToken property.
	//
	// Specify an array of string values to match this event if the actual value of clientRequestToken is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClientRequestToken *[]*string `field:"optional" json:"clientRequestToken" yaml:"clientRequestToken"`
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
	// lastActivityDate property.
	//
	// Specify an array of string values to match this event if the actual value of lastActivityDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastActivityDate *[]*string `field:"optional" json:"lastActivityDate" yaml:"lastActivityDate"`
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
	// pullRequestTargets property.
	//
	// Specify an array of string values to match this event if the actual value of pullRequestTargets is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PullRequestTargets *[]*AWSAPICallViaCloudTrail_PullRequestItem `field:"optional" json:"pullRequestTargets" yaml:"pullRequestTargets"`
	// revisionId property.
	//
	// Specify an array of string values to match this event if the actual value of revisionId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RevisionId *[]*string `field:"optional" json:"revisionId" yaml:"revisionId"`
	// title property.
	//
	// Specify an array of string values to match this event if the actual value of title is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Title *[]*string `field:"optional" json:"title" yaml:"title"`
}


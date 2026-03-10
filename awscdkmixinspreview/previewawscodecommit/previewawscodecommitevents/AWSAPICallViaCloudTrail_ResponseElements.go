package previewawscodecommitevents


// Type definition for ResponseElements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var approvalRules interface{}
//
//   responseElements := &ResponseElements{
//   	AfterBlobId: []*string{
//   		jsii.String("afterBlobId"),
//   	},
//   	AfterCommitId: []*string{
//   		jsii.String("afterCommitId"),
//   	},
//   	BeforeBlobId: []*string{
//   		jsii.String("beforeBlobId"),
//   	},
//   	BeforeCommitId: []*string{
//   		jsii.String("beforeCommitId"),
//   	},
//   	BlobId: []*string{
//   		jsii.String("blobId"),
//   	},
//   	Comment: &Comment{
//   		AuthorArn: []*string{
//   			jsii.String("authorArn"),
//   		},
//   		ClientRequestToken: []*string{
//   			jsii.String("clientRequestToken"),
//   		},
//   		CommentId: []*string{
//   			jsii.String("commentId"),
//   		},
//   		Content: []*string{
//   			jsii.String("content"),
//   		},
//   		CreationDate: []*string{
//   			jsii.String("creationDate"),
//   		},
//   		Deleted: []*string{
//   			jsii.String("deleted"),
//   		},
//   		InReplyTo: []*string{
//   			jsii.String("inReplyTo"),
//   		},
//   		LastModifiedDate: []*string{
//   			jsii.String("lastModifiedDate"),
//   		},
//   	},
//   	CommitId: []*string{
//   		jsii.String("commitId"),
//   	},
//   	DeletedBranch: &DeletedBranch{
//   		BranchName: []*string{
//   			jsii.String("branchName"),
//   		},
//   		CommitId: []*string{
//   			jsii.String("commitId"),
//   		},
//   	},
//   	FilePath: []*string{
//   		jsii.String("filePath"),
//   	},
//   	FilesAdded: []ResponseElementsItem{
//   		&ResponseElementsItem{
//   			AbsolutePath: []*string{
//   				jsii.String("absolutePath"),
//   			},
//   			BlobId: []*string{
//   				jsii.String("blobId"),
//   			},
//   			FileMode: []*string{
//   				jsii.String("fileMode"),
//   			},
//   		},
//   	},
//   	FilesDeleted: []ResponseElementsItem{
//   		&ResponseElementsItem{
//   			AbsolutePath: []*string{
//   				jsii.String("absolutePath"),
//   			},
//   			BlobId: []*string{
//   				jsii.String("blobId"),
//   			},
//   			FileMode: []*string{
//   				jsii.String("fileMode"),
//   			},
//   		},
//   	},
//   	FilesUpdated: []ResponseElementsItem{
//   		&ResponseElementsItem{
//   			AbsolutePath: []*string{
//   				jsii.String("absolutePath"),
//   			},
//   			BlobId: []*string{
//   				jsii.String("blobId"),
//   			},
//   			FileMode: []*string{
//   				jsii.String("fileMode"),
//   			},
//   		},
//   	},
//   	Location: &Location{
//   		FilePath: []*string{
//   			jsii.String("filePath"),
//   		},
//   		FilePosition: []*string{
//   			jsii.String("filePosition"),
//   		},
//   		RelativeFileVersion: []*string{
//   			jsii.String("relativeFileVersion"),
//   		},
//   	},
//   	PullRequest: &PullRequest{
//   		ApprovalRules: []interface{}{
//   			approvalRules,
//   		},
//   		AuthorArn: []*string{
//   			jsii.String("authorArn"),
//   		},
//   		ClientRequestToken: []*string{
//   			jsii.String("clientRequestToken"),
//   		},
//   		CreationDate: []*string{
//   			jsii.String("creationDate"),
//   		},
//   		Description: []*string{
//   			jsii.String("description"),
//   		},
//   		LastActivityDate: []*string{
//   			jsii.String("lastActivityDate"),
//   		},
//   		PullRequestId: []*string{
//   			jsii.String("pullRequestId"),
//   		},
//   		PullRequestStatus: []*string{
//   			jsii.String("pullRequestStatus"),
//   		},
//   		PullRequestTargets: []PullRequestItem{
//   			&PullRequestItem{
//   				DestinationCommit: []*string{
//   					jsii.String("destinationCommit"),
//   				},
//   				DestinationReference: []*string{
//   					jsii.String("destinationReference"),
//   				},
//   				MergeBase: []*string{
//   					jsii.String("mergeBase"),
//   				},
//   				MergeMetadata: &MergeMetadata{
//   					IsMerged: []*string{
//   						jsii.String("isMerged"),
//   					},
//   					MergeCommitId: []*string{
//   						jsii.String("mergeCommitId"),
//   					},
//   					MergedBy: []*string{
//   						jsii.String("mergedBy"),
//   					},
//   					MergeOption: []*string{
//   						jsii.String("mergeOption"),
//   					},
//   				},
//   				RepositoryName: []*string{
//   					jsii.String("repositoryName"),
//   				},
//   				SourceCommit: []*string{
//   					jsii.String("sourceCommit"),
//   				},
//   				SourceReference: []*string{
//   					jsii.String("sourceReference"),
//   				},
//   			},
//   		},
//   		RevisionId: []*string{
//   			jsii.String("revisionId"),
//   		},
//   		Title: []*string{
//   			jsii.String("title"),
//   		},
//   	},
//   	PullRequestId: []*string{
//   		jsii.String("pullRequestId"),
//   	},
//   	RepositoryId: []*string{
//   		jsii.String("repositoryId"),
//   	},
//   	RepositoryMetadata: &RepositoryMetadata{
//   		AccountId: []*string{
//   			jsii.String("accountId"),
//   		},
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		CloneUrlHttp: []*string{
//   			jsii.String("cloneUrlHttp"),
//   		},
//   		CloneUrlSsh: []*string{
//   			jsii.String("cloneUrlSsh"),
//   		},
//   		CreationDate: []*string{
//   			jsii.String("creationDate"),
//   		},
//   		LastModifiedDate: []*string{
//   			jsii.String("lastModifiedDate"),
//   		},
//   		RepositoryDescription: []*string{
//   			jsii.String("repositoryDescription"),
//   		},
//   		RepositoryId: []*string{
//   			jsii.String("repositoryId"),
//   		},
//   		RepositoryName: []*string{
//   			jsii.String("repositoryName"),
//   		},
//   	},
//   	RepositoryName: []*string{
//   		jsii.String("repositoryName"),
//   	},
//   	TreeId: []*string{
//   		jsii.String("treeId"),
//   	},
//   	UploadId: []*string{
//   		jsii.String("uploadId"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_ResponseElements struct {
	// afterBlobId property.
	//
	// Specify an array of string values to match this event if the actual value of afterBlobId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AfterBlobId *[]*string `field:"optional" json:"afterBlobId" yaml:"afterBlobId"`
	// afterCommitId property.
	//
	// Specify an array of string values to match this event if the actual value of afterCommitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AfterCommitId *[]*string `field:"optional" json:"afterCommitId" yaml:"afterCommitId"`
	// beforeBlobId property.
	//
	// Specify an array of string values to match this event if the actual value of beforeBlobId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BeforeBlobId *[]*string `field:"optional" json:"beforeBlobId" yaml:"beforeBlobId"`
	// beforeCommitId property.
	//
	// Specify an array of string values to match this event if the actual value of beforeCommitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BeforeCommitId *[]*string `field:"optional" json:"beforeCommitId" yaml:"beforeCommitId"`
	// blobId property.
	//
	// Specify an array of string values to match this event if the actual value of blobId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BlobId *[]*string `field:"optional" json:"blobId" yaml:"blobId"`
	// comment property.
	//
	// Specify an array of string values to match this event if the actual value of comment is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Comment *AWSAPICallViaCloudTrail_Comment `field:"optional" json:"comment" yaml:"comment"`
	// commitId property.
	//
	// Specify an array of string values to match this event if the actual value of commitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CommitId *[]*string `field:"optional" json:"commitId" yaml:"commitId"`
	// deletedBranch property.
	//
	// Specify an array of string values to match this event if the actual value of deletedBranch is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeletedBranch *AWSAPICallViaCloudTrail_DeletedBranch `field:"optional" json:"deletedBranch" yaml:"deletedBranch"`
	// filePath property.
	//
	// Specify an array of string values to match this event if the actual value of filePath is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FilePath *[]*string `field:"optional" json:"filePath" yaml:"filePath"`
	// filesAdded property.
	//
	// Specify an array of string values to match this event if the actual value of filesAdded is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FilesAdded *[]*AWSAPICallViaCloudTrail_ResponseElementsItem `field:"optional" json:"filesAdded" yaml:"filesAdded"`
	// filesDeleted property.
	//
	// Specify an array of string values to match this event if the actual value of filesDeleted is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FilesDeleted *[]*AWSAPICallViaCloudTrail_ResponseElementsItem `field:"optional" json:"filesDeleted" yaml:"filesDeleted"`
	// filesUpdated property.
	//
	// Specify an array of string values to match this event if the actual value of filesUpdated is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FilesUpdated *[]*AWSAPICallViaCloudTrail_ResponseElementsItem `field:"optional" json:"filesUpdated" yaml:"filesUpdated"`
	// location property.
	//
	// Specify an array of string values to match this event if the actual value of location is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Location *AWSAPICallViaCloudTrail_Location `field:"optional" json:"location" yaml:"location"`
	// pullRequest property.
	//
	// Specify an array of string values to match this event if the actual value of pullRequest is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PullRequest *AWSAPICallViaCloudTrail_PullRequest `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// pullRequestId property.
	//
	// Specify an array of string values to match this event if the actual value of pullRequestId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PullRequestId *[]*string `field:"optional" json:"pullRequestId" yaml:"pullRequestId"`
	// repositoryId property.
	//
	// Specify an array of string values to match this event if the actual value of repositoryId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RepositoryId *[]*string `field:"optional" json:"repositoryId" yaml:"repositoryId"`
	// repositoryMetadata property.
	//
	// Specify an array of string values to match this event if the actual value of repositoryMetadata is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RepositoryMetadata *AWSAPICallViaCloudTrail_RepositoryMetadata `field:"optional" json:"repositoryMetadata" yaml:"repositoryMetadata"`
	// repositoryName property.
	//
	// Specify an array of string values to match this event if the actual value of repositoryName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RepositoryName *[]*string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// treeId property.
	//
	// Specify an array of string values to match this event if the actual value of treeId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TreeId *[]*string `field:"optional" json:"treeId" yaml:"treeId"`
	// uploadId property.
	//
	// Specify an array of string values to match this event if the actual value of uploadId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UploadId *[]*string `field:"optional" json:"uploadId" yaml:"uploadId"`
}


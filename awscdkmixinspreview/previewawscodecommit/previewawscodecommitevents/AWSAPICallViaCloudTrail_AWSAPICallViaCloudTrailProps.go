package previewawscodecommitevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.codecommit@AWSAPICallViaCloudTrail event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var approvalRules interface{}
//
//   aWSAPICallViaCloudTrailProps := &AWSAPICallViaCloudTrailProps{
//   	AdditionalEventData: &AdditionalEventData{
//   		Capabilities: []*string{
//   			jsii.String("capabilities"),
//   		},
//   		Clone: []*string{
//   			jsii.String("clone"),
//   		},
//   		DataTransferred: []*string{
//   			jsii.String("dataTransferred"),
//   		},
//   		Protocol: []*string{
//   			jsii.String("protocol"),
//   		},
//   		RepositoryId: []*string{
//   			jsii.String("repositoryId"),
//   		},
//   		RepositoryName: []*string{
//   			jsii.String("repositoryName"),
//   		},
//   		Shallow: []*string{
//   			jsii.String("shallow"),
//   		},
//   	},
//   	ApiVersion: []*string{
//   		jsii.String("apiVersion"),
//   	},
//   	AwsRegion: []*string{
//   		jsii.String("awsRegion"),
//   	},
//   	ErrorCode: []*string{
//   		jsii.String("errorCode"),
//   	},
//   	ErrorMessage: []*string{
//   		jsii.String("errorMessage"),
//   	},
//   	EventId: []*string{
//   		jsii.String("eventId"),
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
//   	EventName: []*string{
//   		jsii.String("eventName"),
//   	},
//   	EventSource: []*string{
//   		jsii.String("eventSource"),
//   	},
//   	EventTime: []*string{
//   		jsii.String("eventTime"),
//   	},
//   	EventType: []*string{
//   		jsii.String("eventType"),
//   	},
//   	EventVersion: []*string{
//   		jsii.String("eventVersion"),
//   	},
//   	ReadOnly: []*string{
//   		jsii.String("readOnly"),
//   	},
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	RequestParameters: &RequestParameters{
//   		AfterCommitId: []*string{
//   			jsii.String("afterCommitId"),
//   		},
//   		ApprovalRuleTemplateContent: []*string{
//   			jsii.String("approvalRuleTemplateContent"),
//   		},
//   		ApprovalRuleTemplateDescription: []*string{
//   			jsii.String("approvalRuleTemplateDescription"),
//   		},
//   		ApprovalRuleTemplateName: []*string{
//   			jsii.String("approvalRuleTemplateName"),
//   		},
//   		ApprovalState: []*string{
//   			jsii.String("approvalState"),
//   		},
//   		ArchiveType: []*string{
//   			jsii.String("archiveType"),
//   		},
//   		BeforeCommitId: []*string{
//   			jsii.String("beforeCommitId"),
//   		},
//   		BranchName: []*string{
//   			jsii.String("branchName"),
//   		},
//   		ClientRequestToken: []*string{
//   			jsii.String("clientRequestToken"),
//   		},
//   		CommentId: []*string{
//   			jsii.String("commentId"),
//   		},
//   		CommitId: []*string{
//   			jsii.String("commitId"),
//   		},
//   		CommitIds: []*string{
//   			jsii.String("commitIds"),
//   		},
//   		CommitMessage: []*string{
//   			jsii.String("commitMessage"),
//   		},
//   		ConflictDetailLevel: []*string{
//   			jsii.String("conflictDetailLevel"),
//   		},
//   		ConflictResolutionStrategy: []*string{
//   			jsii.String("conflictResolutionStrategy"),
//   		},
//   		Content: []*string{
//   			jsii.String("content"),
//   		},
//   		DefaultBranchName: []*string{
//   			jsii.String("defaultBranchName"),
//   		},
//   		DeleteFiles: []RequestParametersItem{
//   			&RequestParametersItem{
//   				FilePath: []*string{
//   					jsii.String("filePath"),
//   				},
//   			},
//   		},
//   		Description: []*string{
//   			jsii.String("description"),
//   		},
//   		DestinationCommitSpecifier: []*string{
//   			jsii.String("destinationCommitSpecifier"),
//   		},
//   		FileMode: []*string{
//   			jsii.String("fileMode"),
//   		},
//   		FilePath: []*string{
//   			jsii.String("filePath"),
//   		},
//   		FilePaths: []*string{
//   			jsii.String("filePaths"),
//   		},
//   		InReplyTo: []*string{
//   			jsii.String("inReplyTo"),
//   		},
//   		KeepEmptyFolders: []*string{
//   			jsii.String("keepEmptyFolders"),
//   		},
//   		Location: &Location{
//   			FilePath: []*string{
//   				jsii.String("filePath"),
//   			},
//   			FilePosition: []*string{
//   				jsii.String("filePosition"),
//   			},
//   			RelativeFileVersion: []*string{
//   				jsii.String("relativeFileVersion"),
//   			},
//   		},
//   		MaxConflictFiles: []*string{
//   			jsii.String("maxConflictFiles"),
//   		},
//   		MaxMergeHunks: []*string{
//   			jsii.String("maxMergeHunks"),
//   		},
//   		MergeOption: []*string{
//   			jsii.String("mergeOption"),
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   		NewName: []*string{
//   			jsii.String("newName"),
//   		},
//   		OldName: []*string{
//   			jsii.String("oldName"),
//   		},
//   		ParentCommitId: []*string{
//   			jsii.String("parentCommitId"),
//   		},
//   		PullRequestId: []*string{
//   			jsii.String("pullRequestId"),
//   		},
//   		PullRequestIds: []*string{
//   			jsii.String("pullRequestIds"),
//   		},
//   		PullRequestStatus: []*string{
//   			jsii.String("pullRequestStatus"),
//   		},
//   		PutFiles: []RequestParametersItem{
//   			&RequestParametersItem{
//   				FilePath: []*string{
//   					jsii.String("filePath"),
//   				},
//   			},
//   		},
//   		References: []RequestParametersItem1{
//   			&RequestParametersItem1{
//   				Commit: []*string{
//   					jsii.String("commit"),
//   				},
//   				Ref: []*string{
//   					jsii.String("ref"),
//   				},
//   			},
//   		},
//   		RepositoryDescription: []*string{
//   			jsii.String("repositoryDescription"),
//   		},
//   		RepositoryName: []*string{
//   			jsii.String("repositoryName"),
//   		},
//   		ResourceArn: []*string{
//   			jsii.String("resourceArn"),
//   		},
//   		RevisionId: []*string{
//   			jsii.String("revisionId"),
//   		},
//   		S3Bucket: []*string{
//   			jsii.String("s3Bucket"),
//   		},
//   		S3Key: []*string{
//   			jsii.String("s3Key"),
//   		},
//   		SourceCommitSpecifier: []*string{
//   			jsii.String("sourceCommitSpecifier"),
//   		},
//   		TagKeys: []*string{
//   			jsii.String("tagKeys"),
//   		},
//   		Tags: map[string]*string{
//   			"tagsKey": jsii.String("tags"),
//   		},
//   		TargetBranch: []*string{
//   			jsii.String("targetBranch"),
//   		},
//   		Targets: []RequestParametersItem2{
//   			&RequestParametersItem2{
//   				DestinationReference: []*string{
//   					jsii.String("destinationReference"),
//   				},
//   				RepositoryName: []*string{
//   					jsii.String("repositoryName"),
//   				},
//   				SourceReference: []*string{
//   					jsii.String("sourceReference"),
//   				},
//   			},
//   		},
//   		ThrowIfClosed: []*string{
//   			jsii.String("throwIfClosed"),
//   		},
//   		Title: []*string{
//   			jsii.String("title"),
//   		},
//   		UploadId: []*string{
//   			jsii.String("uploadId"),
//   		},
//   	},
//   	Resources: []AwsapiCallViaCloudTrailItem{
//   		&AwsapiCallViaCloudTrailItem{
//   			AccountId: []*string{
//   				jsii.String("accountId"),
//   			},
//   			Arn: []*string{
//   				jsii.String("arn"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   	},
//   	ResponseElements: &ResponseElements{
//   		AfterBlobId: []*string{
//   			jsii.String("afterBlobId"),
//   		},
//   		AfterCommitId: []*string{
//   			jsii.String("afterCommitId"),
//   		},
//   		BeforeBlobId: []*string{
//   			jsii.String("beforeBlobId"),
//   		},
//   		BeforeCommitId: []*string{
//   			jsii.String("beforeCommitId"),
//   		},
//   		BlobId: []*string{
//   			jsii.String("blobId"),
//   		},
//   		Comment: &Comment{
//   			AuthorArn: []*string{
//   				jsii.String("authorArn"),
//   			},
//   			ClientRequestToken: []*string{
//   				jsii.String("clientRequestToken"),
//   			},
//   			CommentId: []*string{
//   				jsii.String("commentId"),
//   			},
//   			Content: []*string{
//   				jsii.String("content"),
//   			},
//   			CreationDate: []*string{
//   				jsii.String("creationDate"),
//   			},
//   			Deleted: []*string{
//   				jsii.String("deleted"),
//   			},
//   			InReplyTo: []*string{
//   				jsii.String("inReplyTo"),
//   			},
//   			LastModifiedDate: []*string{
//   				jsii.String("lastModifiedDate"),
//   			},
//   		},
//   		CommitId: []*string{
//   			jsii.String("commitId"),
//   		},
//   		DeletedBranch: &DeletedBranch{
//   			BranchName: []*string{
//   				jsii.String("branchName"),
//   			},
//   			CommitId: []*string{
//   				jsii.String("commitId"),
//   			},
//   		},
//   		FilePath: []*string{
//   			jsii.String("filePath"),
//   		},
//   		FilesAdded: []ResponseElementsItem{
//   			&ResponseElementsItem{
//   				AbsolutePath: []*string{
//   					jsii.String("absolutePath"),
//   				},
//   				BlobId: []*string{
//   					jsii.String("blobId"),
//   				},
//   				FileMode: []*string{
//   					jsii.String("fileMode"),
//   				},
//   			},
//   		},
//   		FilesDeleted: []ResponseElementsItem{
//   			&ResponseElementsItem{
//   				AbsolutePath: []*string{
//   					jsii.String("absolutePath"),
//   				},
//   				BlobId: []*string{
//   					jsii.String("blobId"),
//   				},
//   				FileMode: []*string{
//   					jsii.String("fileMode"),
//   				},
//   			},
//   		},
//   		FilesUpdated: []ResponseElementsItem{
//   			&ResponseElementsItem{
//   				AbsolutePath: []*string{
//   					jsii.String("absolutePath"),
//   				},
//   				BlobId: []*string{
//   					jsii.String("blobId"),
//   				},
//   				FileMode: []*string{
//   					jsii.String("fileMode"),
//   				},
//   			},
//   		},
//   		Location: &Location{
//   			FilePath: []*string{
//   				jsii.String("filePath"),
//   			},
//   			FilePosition: []*string{
//   				jsii.String("filePosition"),
//   			},
//   			RelativeFileVersion: []*string{
//   				jsii.String("relativeFileVersion"),
//   			},
//   		},
//   		PullRequest: &PullRequest{
//   			ApprovalRules: []interface{}{
//   				approvalRules,
//   			},
//   			AuthorArn: []*string{
//   				jsii.String("authorArn"),
//   			},
//   			ClientRequestToken: []*string{
//   				jsii.String("clientRequestToken"),
//   			},
//   			CreationDate: []*string{
//   				jsii.String("creationDate"),
//   			},
//   			Description: []*string{
//   				jsii.String("description"),
//   			},
//   			LastActivityDate: []*string{
//   				jsii.String("lastActivityDate"),
//   			},
//   			PullRequestId: []*string{
//   				jsii.String("pullRequestId"),
//   			},
//   			PullRequestStatus: []*string{
//   				jsii.String("pullRequestStatus"),
//   			},
//   			PullRequestTargets: []PullRequestItem{
//   				&PullRequestItem{
//   					DestinationCommit: []*string{
//   						jsii.String("destinationCommit"),
//   					},
//   					DestinationReference: []*string{
//   						jsii.String("destinationReference"),
//   					},
//   					MergeBase: []*string{
//   						jsii.String("mergeBase"),
//   					},
//   					MergeMetadata: &MergeMetadata{
//   						IsMerged: []*string{
//   							jsii.String("isMerged"),
//   						},
//   						MergeCommitId: []*string{
//   							jsii.String("mergeCommitId"),
//   						},
//   						MergedBy: []*string{
//   							jsii.String("mergedBy"),
//   						},
//   						MergeOption: []*string{
//   							jsii.String("mergeOption"),
//   						},
//   					},
//   					RepositoryName: []*string{
//   						jsii.String("repositoryName"),
//   					},
//   					SourceCommit: []*string{
//   						jsii.String("sourceCommit"),
//   					},
//   					SourceReference: []*string{
//   						jsii.String("sourceReference"),
//   					},
//   				},
//   			},
//   			RevisionId: []*string{
//   				jsii.String("revisionId"),
//   			},
//   			Title: []*string{
//   				jsii.String("title"),
//   			},
//   		},
//   		PullRequestId: []*string{
//   			jsii.String("pullRequestId"),
//   		},
//   		RepositoryId: []*string{
//   			jsii.String("repositoryId"),
//   		},
//   		RepositoryMetadata: &RepositoryMetadata{
//   			AccountId: []*string{
//   				jsii.String("accountId"),
//   			},
//   			Arn: []*string{
//   				jsii.String("arn"),
//   			},
//   			CloneUrlHttp: []*string{
//   				jsii.String("cloneUrlHttp"),
//   			},
//   			CloneUrlSsh: []*string{
//   				jsii.String("cloneUrlSsh"),
//   			},
//   			CreationDate: []*string{
//   				jsii.String("creationDate"),
//   			},
//   			LastModifiedDate: []*string{
//   				jsii.String("lastModifiedDate"),
//   			},
//   			RepositoryDescription: []*string{
//   				jsii.String("repositoryDescription"),
//   			},
//   			RepositoryId: []*string{
//   				jsii.String("repositoryId"),
//   			},
//   			RepositoryName: []*string{
//   				jsii.String("repositoryName"),
//   			},
//   		},
//   		RepositoryName: []*string{
//   			jsii.String("repositoryName"),
//   		},
//   		TreeId: []*string{
//   			jsii.String("treeId"),
//   		},
//   		UploadId: []*string{
//   			jsii.String("uploadId"),
//   		},
//   	},
//   	SourceIpAddress: []*string{
//   		jsii.String("sourceIpAddress"),
//   	},
//   	UserAgent: []*string{
//   		jsii.String("userAgent"),
//   	},
//   	UserIdentity: &UserIdentity{
//   		AccessKeyId: []*string{
//   			jsii.String("accessKeyId"),
//   		},
//   		AccountId: []*string{
//   			jsii.String("accountId"),
//   		},
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		PrincipalId: []*string{
//   			jsii.String("principalId"),
//   		},
//   		SessionContext: &SessionContext{
//   			Attributes: &Attributes{
//   				CreationDate: []*string{
//   					jsii.String("creationDate"),
//   				},
//   				MfaAuthenticated: []*string{
//   					jsii.String("mfaAuthenticated"),
//   				},
//   			},
//   			SessionIssuer: &SessionIssuer{
//   				AccountId: []*string{
//   					jsii.String("accountId"),
//   				},
//   				Arn: []*string{
//   					jsii.String("arn"),
//   				},
//   				PrincipalId: []*string{
//   					jsii.String("principalId"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   				UserName: []*string{
//   					jsii.String("userName"),
//   				},
//   			},
//   			WebIdFederationData: []*string{
//   				jsii.String("webIdFederationData"),
//   			},
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   		UserName: []*string{
//   			jsii.String("userName"),
//   		},
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps struct {
	// additionalEventData property.
	//
	// Specify an array of string values to match this event if the actual value of additionalEventData is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AdditionalEventData *AWSAPICallViaCloudTrail_AdditionalEventData `field:"optional" json:"additionalEventData" yaml:"additionalEventData"`
	// apiVersion property.
	//
	// Specify an array of string values to match this event if the actual value of apiVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ApiVersion *[]*string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// awsRegion property.
	//
	// Specify an array of string values to match this event if the actual value of awsRegion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsRegion *[]*string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// errorCode property.
	//
	// Specify an array of string values to match this event if the actual value of errorCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorCode *[]*string `field:"optional" json:"errorCode" yaml:"errorCode"`
	// errorMessage property.
	//
	// Specify an array of string values to match this event if the actual value of errorMessage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorMessage *[]*string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// eventID property.
	//
	// Specify an array of string values to match this event if the actual value of eventID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventId *[]*string `field:"optional" json:"eventId" yaml:"eventId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// eventName property.
	//
	// Specify an array of string values to match this event if the actual value of eventName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventName *[]*string `field:"optional" json:"eventName" yaml:"eventName"`
	// eventSource property.
	//
	// Specify an array of string values to match this event if the actual value of eventSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventSource *[]*string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// eventTime property.
	//
	// Specify an array of string values to match this event if the actual value of eventTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventTime *[]*string `field:"optional" json:"eventTime" yaml:"eventTime"`
	// eventType property.
	//
	// Specify an array of string values to match this event if the actual value of eventType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// eventVersion property.
	//
	// Specify an array of string values to match this event if the actual value of eventVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventVersion *[]*string `field:"optional" json:"eventVersion" yaml:"eventVersion"`
	// readOnly property.
	//
	// Specify an array of string values to match this event if the actual value of readOnly is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReadOnly *[]*string `field:"optional" json:"readOnly" yaml:"readOnly"`
	// requestID property.
	//
	// Specify an array of string values to match this event if the actual value of requestID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// requestParameters property.
	//
	// Specify an array of string values to match this event if the actual value of requestParameters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestParameters *AWSAPICallViaCloudTrail_RequestParameters `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// resources property.
	//
	// Specify an array of string values to match this event if the actual value of resources is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Resources *[]*AWSAPICallViaCloudTrail_AwsapiCallViaCloudTrailItem `field:"optional" json:"resources" yaml:"resources"`
	// responseElements property.
	//
	// Specify an array of string values to match this event if the actual value of responseElements is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResponseElements *AWSAPICallViaCloudTrail_ResponseElements `field:"optional" json:"responseElements" yaml:"responseElements"`
	// sourceIPAddress property.
	//
	// Specify an array of string values to match this event if the actual value of sourceIPAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceIpAddress *[]*string `field:"optional" json:"sourceIpAddress" yaml:"sourceIpAddress"`
	// userAgent property.
	//
	// Specify an array of string values to match this event if the actual value of userAgent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserAgent *[]*string `field:"optional" json:"userAgent" yaml:"userAgent"`
	// userIdentity property.
	//
	// Specify an array of string values to match this event if the actual value of userIdentity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserIdentity *AWSAPICallViaCloudTrail_UserIdentity `field:"optional" json:"userIdentity" yaml:"userIdentity"`
}


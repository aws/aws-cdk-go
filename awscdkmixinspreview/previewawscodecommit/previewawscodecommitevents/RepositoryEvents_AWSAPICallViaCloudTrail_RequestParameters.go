package previewawscodecommitevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParameters := &RequestParameters{
//   	AfterCommitId: []*string{
//   		jsii.String("afterCommitId"),
//   	},
//   	ApprovalRuleTemplateContent: []*string{
//   		jsii.String("approvalRuleTemplateContent"),
//   	},
//   	ApprovalRuleTemplateDescription: []*string{
//   		jsii.String("approvalRuleTemplateDescription"),
//   	},
//   	ApprovalRuleTemplateName: []*string{
//   		jsii.String("approvalRuleTemplateName"),
//   	},
//   	ApprovalState: []*string{
//   		jsii.String("approvalState"),
//   	},
//   	ArchiveType: []*string{
//   		jsii.String("archiveType"),
//   	},
//   	BeforeCommitId: []*string{
//   		jsii.String("beforeCommitId"),
//   	},
//   	BranchName: []*string{
//   		jsii.String("branchName"),
//   	},
//   	ClientRequestToken: []*string{
//   		jsii.String("clientRequestToken"),
//   	},
//   	CommentId: []*string{
//   		jsii.String("commentId"),
//   	},
//   	CommitId: []*string{
//   		jsii.String("commitId"),
//   	},
//   	CommitIds: []*string{
//   		jsii.String("commitIds"),
//   	},
//   	CommitMessage: []*string{
//   		jsii.String("commitMessage"),
//   	},
//   	ConflictDetailLevel: []*string{
//   		jsii.String("conflictDetailLevel"),
//   	},
//   	ConflictResolutionStrategy: []*string{
//   		jsii.String("conflictResolutionStrategy"),
//   	},
//   	Content: []*string{
//   		jsii.String("content"),
//   	},
//   	DefaultBranchName: []*string{
//   		jsii.String("defaultBranchName"),
//   	},
//   	DeleteFiles: []RequestParametersItem{
//   		&RequestParametersItem{
//   			FilePath: []*string{
//   				jsii.String("filePath"),
//   			},
//   		},
//   	},
//   	Description: []*string{
//   		jsii.String("description"),
//   	},
//   	DestinationCommitSpecifier: []*string{
//   		jsii.String("destinationCommitSpecifier"),
//   	},
//   	FileMode: []*string{
//   		jsii.String("fileMode"),
//   	},
//   	FilePath: []*string{
//   		jsii.String("filePath"),
//   	},
//   	FilePaths: []*string{
//   		jsii.String("filePaths"),
//   	},
//   	InReplyTo: []*string{
//   		jsii.String("inReplyTo"),
//   	},
//   	KeepEmptyFolders: []*string{
//   		jsii.String("keepEmptyFolders"),
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
//   	MaxConflictFiles: []*string{
//   		jsii.String("maxConflictFiles"),
//   	},
//   	MaxMergeHunks: []*string{
//   		jsii.String("maxMergeHunks"),
//   	},
//   	MergeOption: []*string{
//   		jsii.String("mergeOption"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	NewName: []*string{
//   		jsii.String("newName"),
//   	},
//   	OldName: []*string{
//   		jsii.String("oldName"),
//   	},
//   	ParentCommitId: []*string{
//   		jsii.String("parentCommitId"),
//   	},
//   	PullRequestId: []*string{
//   		jsii.String("pullRequestId"),
//   	},
//   	PullRequestIds: []*string{
//   		jsii.String("pullRequestIds"),
//   	},
//   	PullRequestStatus: []*string{
//   		jsii.String("pullRequestStatus"),
//   	},
//   	PutFiles: []RequestParametersItem{
//   		&RequestParametersItem{
//   			FilePath: []*string{
//   				jsii.String("filePath"),
//   			},
//   		},
//   	},
//   	References: []RequestParametersItem1{
//   		&RequestParametersItem1{
//   			Commit: []*string{
//   				jsii.String("commit"),
//   			},
//   			Ref: []*string{
//   				jsii.String("ref"),
//   			},
//   		},
//   	},
//   	RepositoryDescription: []*string{
//   		jsii.String("repositoryDescription"),
//   	},
//   	RepositoryName: []*string{
//   		jsii.String("repositoryName"),
//   	},
//   	ResourceArn: []*string{
//   		jsii.String("resourceArn"),
//   	},
//   	RevisionId: []*string{
//   		jsii.String("revisionId"),
//   	},
//   	S3Bucket: []*string{
//   		jsii.String("s3Bucket"),
//   	},
//   	S3Key: []*string{
//   		jsii.String("s3Key"),
//   	},
//   	SourceCommitSpecifier: []*string{
//   		jsii.String("sourceCommitSpecifier"),
//   	},
//   	TagKeys: []*string{
//   		jsii.String("tagKeys"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TargetBranch: []*string{
//   		jsii.String("targetBranch"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_AWSAPICallViaCloudTrail_RequestParameters struct {
	// afterCommitId property.
	//
	// Specify an array of string values to match this event if the actual value of afterCommitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AfterCommitId *[]*string `field:"optional" json:"afterCommitId" yaml:"afterCommitId"`
	// approvalRuleTemplateContent property.
	//
	// Specify an array of string values to match this event if the actual value of approvalRuleTemplateContent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ApprovalRuleTemplateContent *[]*string `field:"optional" json:"approvalRuleTemplateContent" yaml:"approvalRuleTemplateContent"`
	// approvalRuleTemplateDescription property.
	//
	// Specify an array of string values to match this event if the actual value of approvalRuleTemplateDescription is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ApprovalRuleTemplateDescription *[]*string `field:"optional" json:"approvalRuleTemplateDescription" yaml:"approvalRuleTemplateDescription"`
	// approvalRuleTemplateName property.
	//
	// Specify an array of string values to match this event if the actual value of approvalRuleTemplateName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ApprovalRuleTemplateName *[]*string `field:"optional" json:"approvalRuleTemplateName" yaml:"approvalRuleTemplateName"`
	// approvalState property.
	//
	// Specify an array of string values to match this event if the actual value of approvalState is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ApprovalState *[]*string `field:"optional" json:"approvalState" yaml:"approvalState"`
	// archiveType property.
	//
	// Specify an array of string values to match this event if the actual value of archiveType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ArchiveType *[]*string `field:"optional" json:"archiveType" yaml:"archiveType"`
	// beforeCommitId property.
	//
	// Specify an array of string values to match this event if the actual value of beforeCommitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BeforeCommitId *[]*string `field:"optional" json:"beforeCommitId" yaml:"beforeCommitId"`
	// branchName property.
	//
	// Specify an array of string values to match this event if the actual value of branchName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BranchName *[]*string `field:"optional" json:"branchName" yaml:"branchName"`
	// clientRequestToken property.
	//
	// Specify an array of string values to match this event if the actual value of clientRequestToken is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClientRequestToken *[]*string `field:"optional" json:"clientRequestToken" yaml:"clientRequestToken"`
	// commentId property.
	//
	// Specify an array of string values to match this event if the actual value of commentId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CommentId *[]*string `field:"optional" json:"commentId" yaml:"commentId"`
	// commitId property.
	//
	// Specify an array of string values to match this event if the actual value of commitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CommitId *[]*string `field:"optional" json:"commitId" yaml:"commitId"`
	// commitIds property.
	//
	// Specify an array of string values to match this event if the actual value of commitIds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CommitIds *[]*string `field:"optional" json:"commitIds" yaml:"commitIds"`
	// commitMessage property.
	//
	// Specify an array of string values to match this event if the actual value of commitMessage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CommitMessage *[]*string `field:"optional" json:"commitMessage" yaml:"commitMessage"`
	// conflictDetailLevel property.
	//
	// Specify an array of string values to match this event if the actual value of conflictDetailLevel is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ConflictDetailLevel *[]*string `field:"optional" json:"conflictDetailLevel" yaml:"conflictDetailLevel"`
	// conflictResolutionStrategy property.
	//
	// Specify an array of string values to match this event if the actual value of conflictResolutionStrategy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ConflictResolutionStrategy *[]*string `field:"optional" json:"conflictResolutionStrategy" yaml:"conflictResolutionStrategy"`
	// content property.
	//
	// Specify an array of string values to match this event if the actual value of content is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Content *[]*string `field:"optional" json:"content" yaml:"content"`
	// defaultBranchName property.
	//
	// Specify an array of string values to match this event if the actual value of defaultBranchName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DefaultBranchName *[]*string `field:"optional" json:"defaultBranchName" yaml:"defaultBranchName"`
	// deleteFiles property.
	//
	// Specify an array of string values to match this event if the actual value of deleteFiles is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeleteFiles *[]*RepositoryEvents_AWSAPICallViaCloudTrail_RequestParametersItem `field:"optional" json:"deleteFiles" yaml:"deleteFiles"`
	// description property.
	//
	// Specify an array of string values to match this event if the actual value of description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// destinationCommitSpecifier property.
	//
	// Specify an array of string values to match this event if the actual value of destinationCommitSpecifier is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DestinationCommitSpecifier *[]*string `field:"optional" json:"destinationCommitSpecifier" yaml:"destinationCommitSpecifier"`
	// fileMode property.
	//
	// Specify an array of string values to match this event if the actual value of fileMode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FileMode *[]*string `field:"optional" json:"fileMode" yaml:"fileMode"`
	// filePath property.
	//
	// Specify an array of string values to match this event if the actual value of filePath is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FilePath *[]*string `field:"optional" json:"filePath" yaml:"filePath"`
	// filePaths property.
	//
	// Specify an array of string values to match this event if the actual value of filePaths is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FilePaths *[]*string `field:"optional" json:"filePaths" yaml:"filePaths"`
	// inReplyTo property.
	//
	// Specify an array of string values to match this event if the actual value of inReplyTo is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InReplyTo *[]*string `field:"optional" json:"inReplyTo" yaml:"inReplyTo"`
	// keepEmptyFolders property.
	//
	// Specify an array of string values to match this event if the actual value of keepEmptyFolders is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	KeepEmptyFolders *[]*string `field:"optional" json:"keepEmptyFolders" yaml:"keepEmptyFolders"`
	// location property.
	//
	// Specify an array of string values to match this event if the actual value of location is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Location *RepositoryEvents_AWSAPICallViaCloudTrail_Location `field:"optional" json:"location" yaml:"location"`
	// maxConflictFiles property.
	//
	// Specify an array of string values to match this event if the actual value of maxConflictFiles is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MaxConflictFiles *[]*string `field:"optional" json:"maxConflictFiles" yaml:"maxConflictFiles"`
	// maxMergeHunks property.
	//
	// Specify an array of string values to match this event if the actual value of maxMergeHunks is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MaxMergeHunks *[]*string `field:"optional" json:"maxMergeHunks" yaml:"maxMergeHunks"`
	// mergeOption property.
	//
	// Specify an array of string values to match this event if the actual value of mergeOption is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MergeOption *[]*string `field:"optional" json:"mergeOption" yaml:"mergeOption"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// newName property.
	//
	// Specify an array of string values to match this event if the actual value of newName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NewName *[]*string `field:"optional" json:"newName" yaml:"newName"`
	// oldName property.
	//
	// Specify an array of string values to match this event if the actual value of oldName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OldName *[]*string `field:"optional" json:"oldName" yaml:"oldName"`
	// parentCommitId property.
	//
	// Specify an array of string values to match this event if the actual value of parentCommitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ParentCommitId *[]*string `field:"optional" json:"parentCommitId" yaml:"parentCommitId"`
	// pullRequestId property.
	//
	// Specify an array of string values to match this event if the actual value of pullRequestId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PullRequestId *[]*string `field:"optional" json:"pullRequestId" yaml:"pullRequestId"`
	// pullRequestIds property.
	//
	// Specify an array of string values to match this event if the actual value of pullRequestIds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PullRequestIds *[]*string `field:"optional" json:"pullRequestIds" yaml:"pullRequestIds"`
	// pullRequestStatus property.
	//
	// Specify an array of string values to match this event if the actual value of pullRequestStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PullRequestStatus *[]*string `field:"optional" json:"pullRequestStatus" yaml:"pullRequestStatus"`
	// putFiles property.
	//
	// Specify an array of string values to match this event if the actual value of putFiles is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PutFiles *[]*RepositoryEvents_AWSAPICallViaCloudTrail_RequestParametersItem `field:"optional" json:"putFiles" yaml:"putFiles"`
	// references property.
	//
	// Specify an array of string values to match this event if the actual value of references is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	References *[]*RepositoryEvents_AWSAPICallViaCloudTrail_RequestParametersItem1 `field:"optional" json:"references" yaml:"references"`
	// repositoryDescription property.
	//
	// Specify an array of string values to match this event if the actual value of repositoryDescription is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RepositoryDescription *[]*string `field:"optional" json:"repositoryDescription" yaml:"repositoryDescription"`
	// repositoryName property.
	//
	// Specify an array of string values to match this event if the actual value of repositoryName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RepositoryName *[]*string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// resourceArn property.
	//
	// Specify an array of string values to match this event if the actual value of resourceArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceArn *[]*string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// revisionId property.
	//
	// Specify an array of string values to match this event if the actual value of revisionId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RevisionId *[]*string `field:"optional" json:"revisionId" yaml:"revisionId"`
	// s3Bucket property.
	//
	// Specify an array of string values to match this event if the actual value of s3Bucket is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3Bucket *[]*string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// s3Key property.
	//
	// Specify an array of string values to match this event if the actual value of s3Key is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3Key *[]*string `field:"optional" json:"s3Key" yaml:"s3Key"`
	// sourceCommitSpecifier property.
	//
	// Specify an array of string values to match this event if the actual value of sourceCommitSpecifier is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceCommitSpecifier *[]*string `field:"optional" json:"sourceCommitSpecifier" yaml:"sourceCommitSpecifier"`
	// tagKeys property.
	//
	// Specify an array of string values to match this event if the actual value of tagKeys is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TagKeys *[]*string `field:"optional" json:"tagKeys" yaml:"tagKeys"`
	// tags property.
	//
	// Specify an array of string values to match this event if the actual value of tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// targetBranch property.
	//
	// Specify an array of string values to match this event if the actual value of targetBranch is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetBranch *[]*string `field:"optional" json:"targetBranch" yaml:"targetBranch"`
}


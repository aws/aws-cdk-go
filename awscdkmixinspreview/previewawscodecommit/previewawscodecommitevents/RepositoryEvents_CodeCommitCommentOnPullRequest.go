package previewawscodecommitevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.codecommit@CodeCommitCommentOnPullRequest event types for Repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeCommitCommentOnPullRequest := #error#.NewCodeCommitCommentOnPullRequest()
//
// Experimental.
type RepositoryEvents_CodeCommitCommentOnPullRequest interface {
}

// The jsii proxy struct for RepositoryEvents_CodeCommitCommentOnPullRequest
type jsiiProxy_RepositoryEvents_CodeCommitCommentOnPullRequest struct {
	_ byte // padding
}

// Experimental.
func NewRepositoryEvents_CodeCommitCommentOnPullRequest() RepositoryEvents_CodeCommitCommentOnPullRequest {
	_init_.Initialize()

	j := jsiiProxy_RepositoryEvents_CodeCommitCommentOnPullRequest{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.CodeCommitCommentOnPullRequest",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRepositoryEvents_CodeCommitCommentOnPullRequest_Override(r RepositoryEvents_CodeCommitCommentOnPullRequest) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.CodeCommitCommentOnPullRequest",
		nil, // no parameters
		r,
	)
}


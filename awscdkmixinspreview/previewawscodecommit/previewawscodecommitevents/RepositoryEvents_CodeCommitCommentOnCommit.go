package previewawscodecommitevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.codecommit@CodeCommitCommentOnCommit event types for Repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeCommitCommentOnCommit := #error#.NewCodeCommitCommentOnCommit()
//
// Experimental.
type RepositoryEvents_CodeCommitCommentOnCommit interface {
}

// The jsii proxy struct for RepositoryEvents_CodeCommitCommentOnCommit
type jsiiProxy_RepositoryEvents_CodeCommitCommentOnCommit struct {
	_ byte // padding
}

// Experimental.
func NewRepositoryEvents_CodeCommitCommentOnCommit() RepositoryEvents_CodeCommitCommentOnCommit {
	_init_.Initialize()

	j := jsiiProxy_RepositoryEvents_CodeCommitCommentOnCommit{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.CodeCommitCommentOnCommit",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRepositoryEvents_CodeCommitCommentOnCommit_Override(r RepositoryEvents_CodeCommitCommentOnCommit) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.CodeCommitCommentOnCommit",
		nil, // no parameters
		r,
	)
}


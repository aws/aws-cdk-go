package previewawscodecommitevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.codecommit@CodeCommitRepositoryStateChange event types for Repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeCommitRepositoryStateChange := #error#.NewCodeCommitRepositoryStateChange()
//
// Experimental.
type RepositoryEvents_CodeCommitRepositoryStateChange interface {
}

// The jsii proxy struct for RepositoryEvents_CodeCommitRepositoryStateChange
type jsiiProxy_RepositoryEvents_CodeCommitRepositoryStateChange struct {
	_ byte // padding
}

// Experimental.
func NewRepositoryEvents_CodeCommitRepositoryStateChange() RepositoryEvents_CodeCommitRepositoryStateChange {
	_init_.Initialize()

	j := jsiiProxy_RepositoryEvents_CodeCommitRepositoryStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.CodeCommitRepositoryStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRepositoryEvents_CodeCommitRepositoryStateChange_Override(r RepositoryEvents_CodeCommitRepositoryStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.CodeCommitRepositoryStateChange",
		nil, // no parameters
		r,
	)
}


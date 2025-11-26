package previewawscodecommitevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodecommit"
)

// EventBridge event patterns for Repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var repositoryRef IRepositoryRef
//
//   repositoryEvents := awscdkmixinspreview.Events.RepositoryEvents_FromRepository(repositoryRef)
//
// Experimental.
type RepositoryEvents interface {
	// EventBridge event pattern for Repository CodeCommit Comment on Commit.
	// Experimental.
	CodeCommitCommentOnCommitPattern(options *RepositoryEvents_CodeCommitCommentOnCommit_CodeCommitCommentOnCommitProps) *awsevents.EventPattern
	// EventBridge event pattern for Repository CodeCommit Comment on Pull Request.
	// Experimental.
	CodeCommitCommentOnPullRequestPattern(options *RepositoryEvents_CodeCommitCommentOnPullRequest_CodeCommitCommentOnPullRequestProps) *awsevents.EventPattern
	// EventBridge event pattern for Repository CodeCommit Repository State Change.
	// Experimental.
	CodeCommitRepositoryStateChangePattern(options *RepositoryEvents_CodeCommitRepositoryStateChange_CodeCommitRepositoryStateChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for RepositoryEvents
type jsiiProxy_RepositoryEvents struct {
	_ byte // padding
}

// Create RepositoryEvents from a Repository reference.
// Experimental.
func RepositoryEvents_FromRepository(repositoryRef interfacesawscodecommit.IRepositoryRef) RepositoryEvents {
	_init_.Initialize()

	if err := validateRepositoryEvents_FromRepositoryParameters(repositoryRef); err != nil {
		panic(err)
	}
	var returns RepositoryEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents",
		"fromRepository",
		[]interface{}{repositoryRef},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEvents) CodeCommitCommentOnCommitPattern(options *RepositoryEvents_CodeCommitCommentOnCommit_CodeCommitCommentOnCommitProps) *awsevents.EventPattern {
	if err := r.validateCodeCommitCommentOnCommitPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		r,
		"codeCommitCommentOnCommitPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEvents) CodeCommitCommentOnPullRequestPattern(options *RepositoryEvents_CodeCommitCommentOnPullRequest_CodeCommitCommentOnPullRequestProps) *awsevents.EventPattern {
	if err := r.validateCodeCommitCommentOnPullRequestPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		r,
		"codeCommitCommentOnPullRequestPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEvents) CodeCommitRepositoryStateChangePattern(options *RepositoryEvents_CodeCommitRepositoryStateChange_CodeCommitRepositoryStateChangeProps) *awsevents.EventPattern {
	if err := r.validateCodeCommitRepositoryStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		r,
		"codeCommitRepositoryStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


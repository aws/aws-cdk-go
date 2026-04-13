package previewawscodecommitevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.codecommit@CodeCommitCommentOnPullRequest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeCommitCommentOnPullRequest := awscdkmixinspreview.Events.NewCodeCommitCommentOnPullRequest()
//
// Experimental.
type CodeCommitCommentOnPullRequest interface {
}

// The jsii proxy struct for CodeCommitCommentOnPullRequest
type jsiiProxy_CodeCommitCommentOnPullRequest struct {
	_ byte // padding
}

// Experimental.
func NewCodeCommitCommentOnPullRequest() CodeCommitCommentOnPullRequest {
	_init_.Initialize()

	j := jsiiProxy_CodeCommitCommentOnPullRequest{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitCommentOnPullRequest",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCodeCommitCommentOnPullRequest_Override(c CodeCommitCommentOnPullRequest) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitCommentOnPullRequest",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CodeCommit Comment on Pull Request.
// Experimental.
func CodeCommitCommentOnPullRequest_EventPattern(options *CodeCommitCommentOnPullRequest_CodeCommitCommentOnPullRequestProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCodeCommitCommentOnPullRequest_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitCommentOnPullRequest",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


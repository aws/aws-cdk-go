package previewawscodecommitevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.codecommit@CodeCommitPullRequestStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeCommitPullRequestStateChange := awscdkmixinspreview.Events.NewCodeCommitPullRequestStateChange()
//
// Experimental.
type CodeCommitPullRequestStateChange interface {
}

// The jsii proxy struct for CodeCommitPullRequestStateChange
type jsiiProxy_CodeCommitPullRequestStateChange struct {
	_ byte // padding
}

// Experimental.
func NewCodeCommitPullRequestStateChange() CodeCommitPullRequestStateChange {
	_init_.Initialize()

	j := jsiiProxy_CodeCommitPullRequestStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitPullRequestStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCodeCommitPullRequestStateChange_Override(c CodeCommitPullRequestStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitPullRequestStateChange",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CodeCommit Pull Request State Change.
// Experimental.
func CodeCommitPullRequestStateChange_EventPattern(options *CodeCommitPullRequestStateChange_CodeCommitPullRequestStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCodeCommitPullRequestStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitPullRequestStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


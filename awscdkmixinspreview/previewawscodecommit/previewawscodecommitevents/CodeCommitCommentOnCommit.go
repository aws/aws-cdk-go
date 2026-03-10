package previewawscodecommitevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.codecommit@CodeCommitCommentOnCommit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeCommitCommentOnCommit := awscdkmixinspreview.Events.NewCodeCommitCommentOnCommit()
//
// Experimental.
type CodeCommitCommentOnCommit interface {
}

// The jsii proxy struct for CodeCommitCommentOnCommit
type jsiiProxy_CodeCommitCommentOnCommit struct {
	_ byte // padding
}

// Experimental.
func NewCodeCommitCommentOnCommit() CodeCommitCommentOnCommit {
	_init_.Initialize()

	j := jsiiProxy_CodeCommitCommentOnCommit{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitCommentOnCommit",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCodeCommitCommentOnCommit_Override(c CodeCommitCommentOnCommit) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitCommentOnCommit",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CodeCommit Comment on Commit.
// Experimental.
func CodeCommitCommentOnCommit_CodeCommitCommentOnCommitPattern(options *CodeCommitCommentOnCommit_CodeCommitCommentOnCommitProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCodeCommitCommentOnCommit_CodeCommitCommentOnCommitPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitCommentOnCommit",
		"codeCommitCommentOnCommitPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


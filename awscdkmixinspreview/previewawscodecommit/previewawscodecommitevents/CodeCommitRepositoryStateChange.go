package previewawscodecommitevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.codecommit@CodeCommitRepositoryStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeCommitRepositoryStateChange := awscdkmixinspreview.Events.NewCodeCommitRepositoryStateChange()
//
// Experimental.
type CodeCommitRepositoryStateChange interface {
}

// The jsii proxy struct for CodeCommitRepositoryStateChange
type jsiiProxy_CodeCommitRepositoryStateChange struct {
	_ byte // padding
}

// Experimental.
func NewCodeCommitRepositoryStateChange() CodeCommitRepositoryStateChange {
	_init_.Initialize()

	j := jsiiProxy_CodeCommitRepositoryStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitRepositoryStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCodeCommitRepositoryStateChange_Override(c CodeCommitRepositoryStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitRepositoryStateChange",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CodeCommit Repository State Change.
// Experimental.
func CodeCommitRepositoryStateChange_EventPattern(options *CodeCommitRepositoryStateChange_CodeCommitRepositoryStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCodeCommitRepositoryStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitRepositoryStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


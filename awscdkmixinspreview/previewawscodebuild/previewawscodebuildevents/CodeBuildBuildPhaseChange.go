package previewawscodebuildevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.codebuild@CodeBuildBuildPhaseChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeBuildBuildPhaseChange := awscdkmixinspreview.Events.NewCodeBuildBuildPhaseChange()
//
// Experimental.
type CodeBuildBuildPhaseChange interface {
}

// The jsii proxy struct for CodeBuildBuildPhaseChange
type jsiiProxy_CodeBuildBuildPhaseChange struct {
	_ byte // padding
}

// Experimental.
func NewCodeBuildBuildPhaseChange() CodeBuildBuildPhaseChange {
	_init_.Initialize()

	j := jsiiProxy_CodeBuildBuildPhaseChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codebuild.events.CodeBuildBuildPhaseChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCodeBuildBuildPhaseChange_Override(c CodeBuildBuildPhaseChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codebuild.events.CodeBuildBuildPhaseChange",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CodeBuild Build Phase Change.
// Experimental.
func CodeBuildBuildPhaseChange_CodeBuildBuildPhaseChangePattern(options *CodeBuildBuildPhaseChange_CodeBuildBuildPhaseChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCodeBuildBuildPhaseChange_CodeBuildBuildPhaseChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codebuild.events.CodeBuildBuildPhaseChange",
		"codeBuildBuildPhaseChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


package previewawscodebuildevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.codebuild@CodeBuildBuildPhaseChange event types for Project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeBuildBuildPhaseChange := #error#.NewCodeBuildBuildPhaseChange()
//
// Experimental.
type ProjectEvents_CodeBuildBuildPhaseChange interface {
}

// The jsii proxy struct for ProjectEvents_CodeBuildBuildPhaseChange
type jsiiProxy_ProjectEvents_CodeBuildBuildPhaseChange struct {
	_ byte // padding
}

// Experimental.
func NewProjectEvents_CodeBuildBuildPhaseChange() ProjectEvents_CodeBuildBuildPhaseChange {
	_init_.Initialize()

	j := jsiiProxy_ProjectEvents_CodeBuildBuildPhaseChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewProjectEvents_CodeBuildBuildPhaseChange_Override(p ProjectEvents_CodeBuildBuildPhaseChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange",
		nil, // no parameters
		p,
	)
}


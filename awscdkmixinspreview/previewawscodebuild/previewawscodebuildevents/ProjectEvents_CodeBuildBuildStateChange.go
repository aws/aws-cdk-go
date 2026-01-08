package previewawscodebuildevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.codebuild@CodeBuildBuildStateChange event types for Project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeBuildBuildStateChange := #error#.NewCodeBuildBuildStateChange()
//
// Experimental.
type ProjectEvents_CodeBuildBuildStateChange interface {
}

// The jsii proxy struct for ProjectEvents_CodeBuildBuildStateChange
type jsiiProxy_ProjectEvents_CodeBuildBuildStateChange struct {
	_ byte // padding
}

// Experimental.
func NewProjectEvents_CodeBuildBuildStateChange() ProjectEvents_CodeBuildBuildStateChange {
	_init_.Initialize()

	j := jsiiProxy_ProjectEvents_CodeBuildBuildStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewProjectEvents_CodeBuildBuildStateChange_Override(p ProjectEvents_CodeBuildBuildStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange",
		nil, // no parameters
		p,
	)
}


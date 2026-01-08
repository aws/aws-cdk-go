package previewawscodebuildevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodebuild"
)

// EventBridge event patterns for Project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var projectRef IProjectRef
//
//   projectEvents := awscdkmixinspreview.Events.ProjectEvents_FromProject(projectRef)
//
// Experimental.
type ProjectEvents interface {
	// EventBridge event pattern for Project CodeBuild Build Phase Change.
	// Experimental.
	CodeBuildBuildPhaseChangePattern(options *ProjectEvents_CodeBuildBuildPhaseChange_CodeBuildBuildPhaseChangeProps) *awsevents.EventPattern
	// EventBridge event pattern for Project CodeBuild Build State Change.
	// Experimental.
	CodeBuildBuildStateChangePattern(options *ProjectEvents_CodeBuildBuildStateChange_CodeBuildBuildStateChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for ProjectEvents
type jsiiProxy_ProjectEvents struct {
	_ byte // padding
}

// Create ProjectEvents from a Project reference.
// Experimental.
func ProjectEvents_FromProject(projectRef interfacesawscodebuild.IProjectRef) ProjectEvents {
	_init_.Initialize()

	if err := validateProjectEvents_FromProjectParameters(projectRef); err != nil {
		panic(err)
	}
	var returns ProjectEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents",
		"fromProject",
		[]interface{}{projectRef},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEvents) CodeBuildBuildPhaseChangePattern(options *ProjectEvents_CodeBuildBuildPhaseChange_CodeBuildBuildPhaseChangeProps) *awsevents.EventPattern {
	if err := p.validateCodeBuildBuildPhaseChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		p,
		"codeBuildBuildPhaseChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectEvents) CodeBuildBuildStateChangePattern(options *ProjectEvents_CodeBuildBuildStateChange_CodeBuildBuildStateChangeProps) *awsevents.EventPattern {
	if err := p.validateCodeBuildBuildStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		p,
		"codeBuildBuildStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


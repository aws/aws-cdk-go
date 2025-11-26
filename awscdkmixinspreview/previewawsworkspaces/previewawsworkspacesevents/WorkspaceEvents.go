package previewawsworkspacesevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsworkspaces"
)

// EventBridge event patterns for Workspace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var workspaceRef IWorkspaceRef
//
//   workspaceEvents := awscdkmixinspreview.Events.WorkspaceEvents_FromWorkspace(workspaceRef)
//
// Experimental.
type WorkspaceEvents interface {
	// EventBridge event pattern for Workspace WorkSpaces Access.
	// Experimental.
	WorkSpacesAccessPattern(options *WorkspaceEvents_WorkSpacesAccess_WorkSpacesAccessProps) *awsevents.EventPattern
}

// The jsii proxy struct for WorkspaceEvents
type jsiiProxy_WorkspaceEvents struct {
	_ byte // padding
}

// Create WorkspaceEvents from a Workspace reference.
// Experimental.
func WorkspaceEvents_FromWorkspace(workspaceRef interfacesawsworkspaces.IWorkspaceRef) WorkspaceEvents {
	_init_.Initialize()

	if err := validateWorkspaceEvents_FromWorkspaceParameters(workspaceRef); err != nil {
		panic(err)
	}
	var returns WorkspaceEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_workspaces.events.WorkspaceEvents",
		"fromWorkspace",
		[]interface{}{workspaceRef},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceEvents) WorkSpacesAccessPattern(options *WorkspaceEvents_WorkSpacesAccess_WorkSpacesAccessProps) *awsevents.EventPattern {
	if err := w.validateWorkSpacesAccessPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		w,
		"workSpacesAccessPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


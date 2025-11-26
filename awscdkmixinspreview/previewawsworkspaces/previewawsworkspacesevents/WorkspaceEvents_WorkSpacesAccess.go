package previewawsworkspacesevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.workspaces@WorkSpacesAccess event types for Workspace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   workSpacesAccess := #error#.NewWorkSpacesAccess()
//
// Experimental.
type WorkspaceEvents_WorkSpacesAccess interface {
}

// The jsii proxy struct for WorkspaceEvents_WorkSpacesAccess
type jsiiProxy_WorkspaceEvents_WorkSpacesAccess struct {
	_ byte // padding
}

// Experimental.
func NewWorkspaceEvents_WorkSpacesAccess() WorkspaceEvents_WorkSpacesAccess {
	_init_.Initialize()

	j := jsiiProxy_WorkspaceEvents_WorkSpacesAccess{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspaces.events.WorkspaceEvents.WorkSpacesAccess",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewWorkspaceEvents_WorkSpacesAccess_Override(w WorkspaceEvents_WorkSpacesAccess) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspaces.events.WorkspaceEvents.WorkSpacesAccess",
		nil, // no parameters
		w,
	)
}


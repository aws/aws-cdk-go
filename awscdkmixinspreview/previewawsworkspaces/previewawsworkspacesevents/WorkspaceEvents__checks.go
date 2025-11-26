//go:build !no_runtime_type_checking

package previewawsworkspacesevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsworkspaces"
)

func (w *jsiiProxy_WorkspaceEvents) validateWorkSpacesAccessPatternParameters(options *WorkspaceEvents_WorkSpacesAccess_WorkSpacesAccessProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateWorkspaceEvents_FromWorkspaceParameters(workspaceRef interfacesawsworkspaces.IWorkspaceRef) error {
	if workspaceRef == nil {
		return fmt.Errorf("parameter workspaceRef is required, but nil was provided")
	}

	return nil
}


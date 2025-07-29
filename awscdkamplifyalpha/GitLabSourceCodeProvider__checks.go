//go:build !no_runtime_type_checking

package awscdkamplifyalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (g *jsiiProxy_GitLabSourceCodeProvider) validateBindParameters(_app App) error {
	if _app == nil {
		return fmt.Errorf("parameter _app is required, but nil was provided")
	}

	return nil
}

func validateNewGitLabSourceCodeProviderParameters(props *GitLabSourceCodeProviderProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


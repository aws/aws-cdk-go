//go:build !no_runtime_type_checking

package awscodebuild

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (l *jsiiProxy_LinuxLambdaBuildImage) validateRunScriptBuildspecParameters(entrypoint *string) error {
	if entrypoint == nil {
		return fmt.Errorf("parameter entrypoint is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LinuxLambdaBuildImage) validateValidateParameters(buildEnvironment *BuildEnvironment) error {
	if buildEnvironment == nil {
		return fmt.Errorf("parameter buildEnvironment is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(buildEnvironment, func() string { return "parameter buildEnvironment" }); err != nil {
		return err
	}

	return nil
}


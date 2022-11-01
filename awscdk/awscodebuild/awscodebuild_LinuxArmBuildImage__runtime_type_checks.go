//go:build !no_runtime_type_checking

package awscodebuild

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsecr"
)

func (l *jsiiProxy_LinuxArmBuildImage) validateRunScriptBuildspecParameters(entrypoint *string) error {
	if entrypoint == nil {
		return fmt.Errorf("parameter entrypoint is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LinuxArmBuildImage) validateValidateParameters(buildEnvironment *BuildEnvironment) error {
	if buildEnvironment == nil {
		return fmt.Errorf("parameter buildEnvironment is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(buildEnvironment, func() string { return "parameter buildEnvironment" }); err != nil {
		return err
	}

	return nil
}

func validateLinuxArmBuildImage_FromCodeBuildImageIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func validateLinuxArmBuildImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	return nil
}


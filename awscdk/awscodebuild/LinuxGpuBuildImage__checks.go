//go:build !no_runtime_type_checking

package awscodebuild

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_LinuxGpuBuildImage) validateBindParameters(scope constructs.Construct, project IProject, _options *BuildImageBindOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if _options == nil {
		return fmt.Errorf("parameter _options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (l *jsiiProxy_LinuxGpuBuildImage) validateRunScriptBuildspecParameters(entrypoint *string) error {
	if entrypoint == nil {
		return fmt.Errorf("parameter entrypoint is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LinuxGpuBuildImage) validateValidateParameters(buildEnvironment *BuildEnvironment) error {
	if buildEnvironment == nil {
		return fmt.Errorf("parameter buildEnvironment is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(buildEnvironment, func() string { return "parameter buildEnvironment" }); err != nil {
		return err
	}

	return nil
}

func validateLinuxGpuBuildImage_AwsDeepLearningContainersImageParameters(repositoryName *string, tag *string) error {
	if repositoryName == nil {
		return fmt.Errorf("parameter repositoryName is required, but nil was provided")
	}

	if tag == nil {
		return fmt.Errorf("parameter tag is required, but nil was provided")
	}

	return nil
}

func validateLinuxGpuBuildImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	return nil
}


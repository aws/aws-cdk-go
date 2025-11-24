//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AgentRuntimeArtifact) validateBindParameters(scope constructs.Construct, runtime Runtime) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if runtime == nil {
		return fmt.Errorf("parameter runtime is required, but nil was provided")
	}

	return nil
}

func validateAgentRuntimeArtifact_FromAssetParameters(directory *string, options *awsecrassets.DockerImageAssetOptions) error {
	if directory == nil {
		return fmt.Errorf("parameter directory is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateAgentRuntimeArtifact_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	return nil
}

func validateAgentRuntimeArtifact_FromS3Parameters(s3Location *awss3.Location, runtime AgentCoreRuntime, entrypoint *[]*string) error {
	if s3Location == nil {
		return fmt.Errorf("parameter s3Location is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(s3Location, func() string { return "parameter s3Location" }); err != nil {
		return err
	}

	if runtime == "" {
		return fmt.Errorf("parameter runtime is required, but nil was provided")
	}

	if entrypoint == nil {
		return fmt.Errorf("parameter entrypoint is required, but nil was provided")
	}

	return nil
}


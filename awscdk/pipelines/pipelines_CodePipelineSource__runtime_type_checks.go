//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscodecommit"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

func (c *jsiiProxy_CodePipelineSource) validateAddDependencyFileSetParameters(fs FileSet) error {
	if fs == nil {
		return fmt.Errorf("parameter fs is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CodePipelineSource) validateAddStepDependencyParameters(step Step) error {
	if step == nil {
		return fmt.Errorf("parameter step is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CodePipelineSource) validateConfigurePrimaryOutputParameters(fs FileSet) error {
	if fs == nil {
		return fmt.Errorf("parameter fs is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CodePipelineSource) validateDiscoverReferencedOutputsParameters(structure interface{}) error {
	if structure == nil {
		return fmt.Errorf("parameter structure is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CodePipelineSource) validateGetActionParameters(output awscodepipeline.Artifact, actionName *string, runOrder *float64) error {
	if output == nil {
		return fmt.Errorf("parameter output is required, but nil was provided")
	}

	if actionName == nil {
		return fmt.Errorf("parameter actionName is required, but nil was provided")
	}

	if runOrder == nil {
		return fmt.Errorf("parameter runOrder is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CodePipelineSource) validateProduceActionParameters(stage awscodepipeline.IStage, options *ProduceActionOptions) error {
	if stage == nil {
		return fmt.Errorf("parameter stage is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CodePipelineSource) validateSourceAttributeParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateCodePipelineSource_CodeCommitParameters(repository awscodecommit.IRepository, branch *string, props *CodeCommitSourceOptions) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	if branch == nil {
		return fmt.Errorf("parameter branch is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateCodePipelineSource_ConnectionParameters(repoString *string, branch *string, props *ConnectionSourceOptions) error {
	if repoString == nil {
		return fmt.Errorf("parameter repoString is required, but nil was provided")
	}

	if branch == nil {
		return fmt.Errorf("parameter branch is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateCodePipelineSource_EcrParameters(repository awsecr.IRepository, props *ECRSourceOptions) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateCodePipelineSource_GitHubParameters(repoString *string, branch *string, props *GitHubSourceOptions) error {
	if repoString == nil {
		return fmt.Errorf("parameter repoString is required, but nil was provided")
	}

	if branch == nil {
		return fmt.Errorf("parameter branch is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateCodePipelineSource_S3Parameters(bucket awss3.IBucket, objectKey *string, props *S3SourceOptions) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if objectKey == nil {
		return fmt.Errorf("parameter objectKey is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateCodePipelineSource_SequenceParameters(steps *[]Step) error {
	if steps == nil {
		return fmt.Errorf("parameter steps is required, but nil was provided")
	}

	return nil
}

func validateNewCodePipelineSourceParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}


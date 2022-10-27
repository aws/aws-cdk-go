//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

func (c *jsiiProxy_ConfirmPermissionsBroadening) validateAddDependencyFileSetParameters(fs FileSet) error {
	if fs == nil {
		return fmt.Errorf("parameter fs is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) validateAddStepDependencyParameters(step Step) error {
	if step == nil {
		return fmt.Errorf("parameter step is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) validateConfigurePrimaryOutputParameters(fs FileSet) error {
	if fs == nil {
		return fmt.Errorf("parameter fs is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) validateDiscoverReferencedOutputsParameters(structure interface{}) error {
	if structure == nil {
		return fmt.Errorf("parameter structure is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) validateProduceActionParameters(stage awscodepipeline.IStage, options *ProduceActionOptions) error {
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

func validateConfirmPermissionsBroadening_SequenceParameters(steps *[]Step) error {
	if steps == nil {
		return fmt.Errorf("parameter steps is required, but nil was provided")
	}

	return nil
}

func validateNewConfirmPermissionsBroadeningParameters(id *string, props *PermissionsBroadeningCheckProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


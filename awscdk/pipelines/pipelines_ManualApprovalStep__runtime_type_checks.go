//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (m *jsiiProxy_ManualApprovalStep) validateAddDependencyFileSetParameters(fs FileSet) error {
	if fs == nil {
		return fmt.Errorf("parameter fs is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManualApprovalStep) validateAddStepDependencyParameters(step Step) error {
	if step == nil {
		return fmt.Errorf("parameter step is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManualApprovalStep) validateConfigurePrimaryOutputParameters(fs FileSet) error {
	if fs == nil {
		return fmt.Errorf("parameter fs is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManualApprovalStep) validateDiscoverReferencedOutputsParameters(structure interface{}) error {
	if structure == nil {
		return fmt.Errorf("parameter structure is required, but nil was provided")
	}

	return nil
}

func validateManualApprovalStep_SequenceParameters(steps *[]Step) error {
	if steps == nil {
		return fmt.Errorf("parameter steps is required, but nil was provided")
	}

	return nil
}

func validateNewManualApprovalStepParameters(id *string, props *ManualApprovalStepProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


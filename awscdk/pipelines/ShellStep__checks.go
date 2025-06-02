//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (s *jsiiProxy_ShellStep) validateAddDependencyFileSetParameters(fs FileSet) error {
	if fs == nil {
		return fmt.Errorf("parameter fs is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ShellStep) validateAddOutputDirectoryParameters(directory *string) error {
	if directory == nil {
		return fmt.Errorf("parameter directory is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ShellStep) validateAddStepDependencyParameters(step Step) error {
	if step == nil {
		return fmt.Errorf("parameter step is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ShellStep) validateConfigurePrimaryOutputParameters(fs FileSet) error {
	if fs == nil {
		return fmt.Errorf("parameter fs is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ShellStep) validateDiscoverReferencedOutputsParameters(structure interface{}) error {
	if structure == nil {
		return fmt.Errorf("parameter structure is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ShellStep) validatePrimaryOutputDirectoryParameters(directory *string) error {
	if directory == nil {
		return fmt.Errorf("parameter directory is required, but nil was provided")
	}

	return nil
}

func validateShellStep_SequenceParameters(steps *[]Step) error {
	if steps == nil {
		return fmt.Errorf("parameter steps is required, but nil was provided")
	}

	return nil
}

func validateNewShellStepParameters(id *string, props *ShellStepProps) error {
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


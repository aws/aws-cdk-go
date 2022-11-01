//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"
)

func (s *jsiiProxy_Step) validateAddDependencyFileSetParameters(fs FileSet) error {
	if fs == nil {
		return fmt.Errorf("parameter fs is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Step) validateAddStepDependencyParameters(step Step) error {
	if step == nil {
		return fmt.Errorf("parameter step is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Step) validateConfigurePrimaryOutputParameters(fs FileSet) error {
	if fs == nil {
		return fmt.Errorf("parameter fs is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Step) validateDiscoverReferencedOutputsParameters(structure interface{}) error {
	if structure == nil {
		return fmt.Errorf("parameter structure is required, but nil was provided")
	}

	return nil
}

func validateStep_SequenceParameters(steps *[]Step) error {
	if steps == nil {
		return fmt.Errorf("parameter steps is required, but nil was provided")
	}

	return nil
}

func validateNewStepParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}


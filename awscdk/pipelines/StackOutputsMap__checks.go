//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"
)

func (s *jsiiProxy_StackOutputsMap) validateToCodePipelineParameters(x StackOutputReference) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewStackOutputsMapParameters(pipeline PipelineBase) error {
	if pipeline == nil {
		return fmt.Errorf("parameter pipeline is required, but nil was provided")
	}

	return nil
}


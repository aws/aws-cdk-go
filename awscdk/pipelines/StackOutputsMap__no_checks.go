//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StackOutputsMap) validateToCodePipelineParameters(x StackOutputReference) error {
	return nil
}

func validateNewStackOutputsMapParameters(pipeline PipelineBase) error {
	return nil
}


//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PipelineBase) validateAddStageParameters(stage awscdk.Stage, options *AddStageOpts) error {
	return nil
}

func (p *jsiiProxy_PipelineBase) validateAddWaveParameters(id *string, options *WaveOptions) error {
	return nil
}

func validatePipelineBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePipelineBase_IsPipelineParameters(x interface{}) error {
	return nil
}

func validateNewPipelineBaseParameters(scope constructs.Construct, id *string, props *PipelineBaseProps) error {
	return nil
}


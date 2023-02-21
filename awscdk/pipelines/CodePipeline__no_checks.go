//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodePipeline) validateAddStageParameters(stage awscdk.Stage, options *AddStageOpts) error {
	return nil
}

func (c *jsiiProxy_CodePipeline) validateAddWaveParameters(id *string, options *WaveOptions) error {
	return nil
}

func validateCodePipeline_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCodePipeline_IsPipelineParameters(x interface{}) error {
	return nil
}

func validateNewCodePipelineParameters(scope constructs.Construct, id *string, props *CodePipelineProps) error {
	return nil
}


//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PipelineInvokeAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineInvokeAction) validateBoundParameters(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineInvokeAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (p *jsiiProxy_PipelineInvokeAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewPipelineInvokeActionParameters(props *PipelineInvokeActionProps) error {
	return nil
}


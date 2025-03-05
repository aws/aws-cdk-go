//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ICodePipelineActionFactory) validateProduceActionParameters(stage awscodepipeline.IStage, options *ProduceActionOptions) error {
	return nil
}


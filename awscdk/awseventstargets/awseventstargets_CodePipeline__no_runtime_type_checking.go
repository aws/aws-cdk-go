//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodePipeline) validateBindParameters(_rule awsevents.IRule) error {
	return nil
}

func validateNewCodePipelineParameters(pipeline awscodepipeline.IPipeline, options *CodePipelineTargetOptions) error {
	return nil
}


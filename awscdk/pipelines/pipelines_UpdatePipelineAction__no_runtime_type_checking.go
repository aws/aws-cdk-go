//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UpdatePipelineAction) validateBindParameters(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (u *jsiiProxy_UpdatePipelineAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (u *jsiiProxy_UpdatePipelineAction) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (u *jsiiProxy_UpdatePipelineAction) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateUpdatePipelineAction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewUpdatePipelineActionParameters(scope constructs.Construct, id *string, props *UpdatePipelineActionProps) error {
	return nil
}


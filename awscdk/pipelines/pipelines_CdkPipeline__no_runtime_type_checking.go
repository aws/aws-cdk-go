//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CdkPipeline) validateAddApplicationStageParameters(appStage awscdk.Stage, options *AddStageOptions) error {
	return nil
}

func (c *jsiiProxy_CdkPipeline) validateAddStageParameters(stageName *string, options *BaseStageOptions) error {
	return nil
}

func (c *jsiiProxy_CdkPipeline) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CdkPipeline) validateStackOutputParameters(cfnOutput awscdk.CfnOutput) error {
	return nil
}

func (c *jsiiProxy_CdkPipeline) validateStageParameters(stageName *string) error {
	return nil
}

func (c *jsiiProxy_CdkPipeline) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateCdkPipeline_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCdkPipelineParameters(scope constructs.Construct, id *string, props *CdkPipelineProps) error {
	return nil
}


//go:build no_runtime_type_checking

package awscdksagemakeralpha

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Pipeline) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateGrantStartPipelineExecutionParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validatePipeline_FromPipelineArnParameters(scope constructs.Construct, id *string, pipelineArn *string) error {
	return nil
}

func validatePipeline_FromPipelineAttributesParameters(scope constructs.Construct, id *string, attrs *PipelineAttributes) error {
	return nil
}

func validatePipeline_FromPipelineNameParameters(scope constructs.Construct, id *string, pipelineName *string) error {
	return nil
}

func validatePipeline_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePipeline_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePipeline_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPipelineParameters(scope constructs.Construct, id *string, props *PipelineProps) error {
	return nil
}


//go:build no_runtime_type_checking

package awscdkpipesalpha

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Pipe) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_Pipe) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_Pipe) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePipe_FromPipeNameParameters(scope constructs.Construct, id *string, pipeName *string) error {
	return nil
}

func validatePipe_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePipe_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePipe_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPipeParameters(scope constructs.Construct, id *string, props *PipeProps) error {
	return nil
}


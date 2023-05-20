//go:build no_runtime_type_checking

package awscdkgameliftalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScriptBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_ScriptBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_ScriptBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateScriptBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateScriptBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateScriptBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewScriptBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}


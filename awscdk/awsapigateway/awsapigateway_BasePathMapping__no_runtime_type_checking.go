//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BasePathMapping) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BasePathMapping) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BasePathMapping) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (b *jsiiProxy_BasePathMapping) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (b *jsiiProxy_BasePathMapping) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateBasePathMapping_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBasePathMapping_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewBasePathMappingParameters(scope constructs.Construct, id *string, props *BasePathMappingProps) error {
	return nil
}


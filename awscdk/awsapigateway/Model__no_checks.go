//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Model) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (m *jsiiProxy_Model) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (m *jsiiProxy_Model) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateModel_FromModelNameParameters(scope constructs.Construct, id *string, modelName *string) error {
	return nil
}

func validateModel_IsConstructParameters(x interface{}) error {
	return nil
}

func validateModel_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateModel_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewModelParameters(scope constructs.Construct, id *string, props *ModelProps) error {
	return nil
}


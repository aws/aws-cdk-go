//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RequestValidator) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RequestValidator) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RequestValidator) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateRequestValidator_FromRequestValidatorIdParameters(scope constructs.Construct, id *string, requestValidatorId *string) error {
	return nil
}

func validateRequestValidator_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRequestValidator_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRequestValidator_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRequestValidatorParameters(scope constructs.Construct, id *string, props *RequestValidatorProps) error {
	return nil
}


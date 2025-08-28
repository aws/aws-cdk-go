//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HealthCheck) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (h *jsiiProxy_HealthCheck) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (h *jsiiProxy_HealthCheck) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateHealthCheck_FromHealthCheckIdParameters(scope constructs.Construct, id *string, healthCheckId *string) error {
	return nil
}

func validateHealthCheck_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHealthCheck_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateHealthCheck_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewHealthCheckParameters(scope constructs.Construct, id *string, props *HealthCheckProps) error {
	return nil
}


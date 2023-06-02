//go:build no_runtime_type_checking

package awscdkapprunneralpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Service) validateAddEnvironmentVariableParameters(name *string, value *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateAddSecretParameters(name *string, secret Secret) error {
	return nil
}

func (s *jsiiProxy_Service) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateService_FromServiceAttributesParameters(scope constructs.Construct, id *string, attrs *ServiceAttributes) error {
	return nil
}

func validateService_FromServiceNameParameters(scope constructs.Construct, id *string, serviceName *string) error {
	return nil
}

func validateService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateService_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateService_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewServiceParameters(scope constructs.Construct, id *string, props *ServiceProps) error {
	return nil
}


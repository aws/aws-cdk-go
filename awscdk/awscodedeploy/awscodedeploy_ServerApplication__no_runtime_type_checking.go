//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServerApplication) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_ServerApplication) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_ServerApplication) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateServerApplication_FromServerApplicationNameParameters(scope constructs.Construct, id *string, serverApplicationName *string) error {
	return nil
}

func validateServerApplication_IsConstructParameters(x interface{}) error {
	return nil
}

func validateServerApplication_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateServerApplication_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewServerApplicationParameters(scope constructs.Construct, id *string, props *ServerApplicationProps) error {
	return nil
}


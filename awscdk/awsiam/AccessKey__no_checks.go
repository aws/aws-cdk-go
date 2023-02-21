//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessKey) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AccessKey) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AccessKey) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAccessKey_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAccessKey_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAccessKey_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAccessKeyParameters(scope constructs.Construct, id *string, props *AccessKeyProps) error {
	return nil
}


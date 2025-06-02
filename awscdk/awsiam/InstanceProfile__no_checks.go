//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InstanceProfile) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_InstanceProfile) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_InstanceProfile) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateInstanceProfile_FromInstanceProfileArnParameters(scope constructs.Construct, id *string, instanceProfileArn *string) error {
	return nil
}

func validateInstanceProfile_FromInstanceProfileAttributesParameters(scope constructs.Construct, id *string, attrs *InstanceProfileAttributes) error {
	return nil
}

func validateInstanceProfile_FromInstanceProfileNameParameters(scope constructs.Construct, id *string, instanceProfileName *string) error {
	return nil
}

func validateInstanceProfile_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInstanceProfile_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateInstanceProfile_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewInstanceProfileParameters(scope constructs.Construct, id *string, props *InstanceProfileProps) error {
	return nil
}


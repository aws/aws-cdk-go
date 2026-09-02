//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InputSecurityGroup) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (i *jsiiProxy_InputSecurityGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_InputSecurityGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_InputSecurityGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateInputSecurityGroup_FromInputSecurityGroupArnParameters(scope constructs.Construct, id *string, inputSecurityGroupArn *string) error {
	return nil
}

func validateInputSecurityGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInputSecurityGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateInputSecurityGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewInputSecurityGroupParameters(scope constructs.Construct, id *string, props *InputSecurityGroupProps) error {
	return nil
}


//go:build no_runtime_type_checking

package awssynthetics

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_Group) validateAddCanaryParameters(canary ICanary) error {
	return nil
}

func (g *jsiiProxy_Group) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (g *jsiiProxy_Group) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_Group) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_Group) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateGroup_FromGroupArnParameters(scope constructs.Construct, id *string, groupArn *string) error {
	return nil
}

func validateGroup_FromGroupNameParameters(scope constructs.Construct, id *string, groupName *string) error {
	return nil
}

func validateGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGroupParameters(scope constructs.Construct, id *string, props *GroupProps) error {
	return nil
}


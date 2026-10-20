//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KnowledgeBaseBase) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBaseBase) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBaseBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBaseBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBaseBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKnowledgeBaseBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKnowledgeBaseBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKnowledgeBaseBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKnowledgeBaseBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}


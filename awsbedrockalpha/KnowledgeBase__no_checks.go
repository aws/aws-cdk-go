//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KnowledgeBase) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBase) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKnowledgeBase_FromKnowledgeBaseArnParameters(scope constructs.Construct, id *string, knowledgeBaseArn *string) error {
	return nil
}

func validateKnowledgeBase_FromKnowledgeBaseAttributesParameters(scope constructs.Construct, id *string, attrs *KnowledgeBaseAttributes) error {
	return nil
}

func validateKnowledgeBase_FromKnowledgeBaseIdParameters(scope constructs.Construct, id *string, knowledgeBaseId *string) error {
	return nil
}

func validateKnowledgeBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKnowledgeBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKnowledgeBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKnowledgeBaseParameters(scope constructs.Construct, id *string, props *KnowledgeBaseProps) error {
	return nil
}


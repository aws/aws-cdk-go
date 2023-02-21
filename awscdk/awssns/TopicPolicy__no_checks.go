//go:build no_runtime_type_checking

package awssns

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TopicPolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TopicPolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TopicPolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTopicPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTopicPolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTopicPolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTopicPolicyParameters(scope constructs.Construct, id *string, props *TopicPolicyProps) error {
	return nil
}


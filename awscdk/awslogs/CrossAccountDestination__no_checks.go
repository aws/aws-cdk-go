//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CrossAccountDestination) validateAddToPolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (c *jsiiProxy_CrossAccountDestination) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CrossAccountDestination) validateBindParameters(_scope constructs.Construct, _sourceLogGroup ILogGroup) error {
	return nil
}

func (c *jsiiProxy_CrossAccountDestination) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CrossAccountDestination) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCrossAccountDestination_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCrossAccountDestination_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCrossAccountDestination_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCrossAccountDestinationParameters(scope constructs.Construct, id *string, props *CrossAccountDestinationProps) error {
	return nil
}


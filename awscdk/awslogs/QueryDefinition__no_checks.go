//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueryDefinition) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (q *jsiiProxy_QueryDefinition) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (q *jsiiProxy_QueryDefinition) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateQueryDefinition_IsConstructParameters(x interface{}) error {
	return nil
}

func validateQueryDefinition_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateQueryDefinition_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewQueryDefinitionParameters(scope constructs.Construct, id *string, props *QueryDefinitionProps) error {
	return nil
}


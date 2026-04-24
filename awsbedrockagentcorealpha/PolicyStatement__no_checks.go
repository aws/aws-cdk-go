//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicyStatement) validateForPrincipalParameters(entityType *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStatement) validateForPrincipalInGroupParameters(groupType *string, groupId *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStatement) validateOnActionParameters(action *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStatement) validateOnActionsParameters(actions *[]*string) error {
	return nil
}

func (p *jsiiProxy_PolicyStatement) validateOnResourceParameters(entityType *string, entityArn *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStatement) validateOnResourceTypeParameters(entityType *string) error {
	return nil
}

func validatePolicyStatement_FromCedarParameters(cedarStatement *string) error {
	return nil
}


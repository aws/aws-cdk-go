//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicyStatement) validateAddAccountConditionParameters(accountId *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStatement) validateAddArnPrincipalParameters(arn *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStatement) validateAddAwsAccountPrincipalParameters(accountId *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStatement) validateAddCanonicalUserPrincipalParameters(canonicalUserId *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStatement) validateAddConditionParameters(key *string, value interface{}) error {
	return nil
}

func (p *jsiiProxy_PolicyStatement) validateAddConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func (p *jsiiProxy_PolicyStatement) validateAddFederatedPrincipalParameters(federated interface{}, conditions *map[string]interface{}) error {
	return nil
}

func (p *jsiiProxy_PolicyStatement) validateAddServicePrincipalParameters(service *string, opts *ServicePrincipalOpts) error {
	return nil
}

func (p *jsiiProxy_PolicyStatement) validateCopyParameters(overrides *PolicyStatementProps) error {
	return nil
}

func validatePolicyStatement_FromJsonParameters(obj interface{}) error {
	return nil
}

func (j *jsiiProxy_PolicyStatement) validateSetEffectParameters(val Effect) error {
	return nil
}

func validateNewPolicyStatementParameters(props *PolicyStatementProps) error {
	return nil
}


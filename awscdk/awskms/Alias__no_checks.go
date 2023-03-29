//go:build no_runtime_type_checking

package awskms

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Alias) validateAddAliasParameters(alias *string) error {
	return nil
}

func (a *jsiiProxy_Alias) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (a *jsiiProxy_Alias) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_Alias) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_Alias) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_Alias) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_Alias) validateGrantDecryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_Alias) validateGrantEncryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_Alias) validateGrantEncryptDecryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateAlias_FromAliasAttributesParameters(scope constructs.Construct, id *string, attrs *AliasAttributes) error {
	return nil
}

func validateAlias_FromAliasNameParameters(scope constructs.Construct, id *string, aliasName *string) error {
	return nil
}

func validateAlias_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAlias_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAlias_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAliasParameters(scope constructs.Construct, id *string, props *AliasProps) error {
	return nil
}


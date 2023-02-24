//go:build no_runtime_type_checking

package awskms

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_Key) validateAddAliasParameters(aliasName *string) error {
	return nil
}

func (k *jsiiProxy_Key) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (k *jsiiProxy_Key) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_Key) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_Key) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (k *jsiiProxy_Key) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_Key) validateGrantAdminParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_Key) validateGrantDecryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_Key) validateGrantEncryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_Key) validateGrantEncryptDecryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateKey_FromCfnKeyParameters(cfnKey CfnKey) error {
	return nil
}

func validateKey_FromKeyArnParameters(scope constructs.Construct, id *string, keyArn *string) error {
	return nil
}

func validateKey_FromLookupParameters(scope constructs.Construct, id *string, options *KeyLookupOptions) error {
	return nil
}

func validateKey_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKey_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKey_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKeyParameters(scope constructs.Construct, id *string, props *KeyProps) error {
	return nil
}


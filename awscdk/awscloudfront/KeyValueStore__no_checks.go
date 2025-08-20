//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyValueStore) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KeyValueStore) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KeyValueStore) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKeyValueStore_FromKeyValueStoreArnParameters(scope constructs.Construct, id *string, keyValueStoreArn *string) error {
	return nil
}

func validateKeyValueStore_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKeyValueStore_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKeyValueStore_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKeyValueStoreParameters(scope constructs.Construct, id *string, props *KeyValueStoreProps) error {
	return nil
}


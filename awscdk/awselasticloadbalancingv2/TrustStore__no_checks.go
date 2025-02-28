//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TrustStore) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TrustStore) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TrustStore) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTrustStore_FromTrustStoreArnParameters(scope constructs.Construct, id *string, trustStoreArn *string) error {
	return nil
}

func validateTrustStore_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTrustStore_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTrustStore_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTrustStoreParameters(scope constructs.Construct, id *string, props *TrustStoreProps) error {
	return nil
}


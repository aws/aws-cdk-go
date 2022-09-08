//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CachePolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CachePolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CachePolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_CachePolicy) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CachePolicy) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateCachePolicy_FromCachePolicyIdParameters(scope constructs.Construct, id *string, cachePolicyId *string) error {
	return nil
}

func validateCachePolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCachePolicy_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewCachePolicyParameters(scope constructs.Construct, id *string, props *CachePolicyProps) error {
	return nil
}


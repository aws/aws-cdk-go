//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResponseHeadersPolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_ResponseHeadersPolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_ResponseHeadersPolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateResponseHeadersPolicy_FromResponseHeadersPolicyIdParameters(scope constructs.Construct, id *string, responseHeadersPolicyId *string) error {
	return nil
}

func validateResponseHeadersPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateResponseHeadersPolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateResponseHeadersPolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewResponseHeadersPolicyParameters(scope constructs.Construct, id *string, props *ResponseHeadersPolicyProps) error {
	return nil
}


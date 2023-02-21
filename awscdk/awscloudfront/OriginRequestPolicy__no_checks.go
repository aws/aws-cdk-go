//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OriginRequestPolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (o *jsiiProxy_OriginRequestPolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (o *jsiiProxy_OriginRequestPolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateOriginRequestPolicy_FromOriginRequestPolicyIdParameters(scope constructs.Construct, id *string, originRequestPolicyId *string) error {
	return nil
}

func validateOriginRequestPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateOriginRequestPolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateOriginRequestPolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewOriginRequestPolicyParameters(scope constructs.Construct, id *string, props *OriginRequestPolicyProps) error {
	return nil
}


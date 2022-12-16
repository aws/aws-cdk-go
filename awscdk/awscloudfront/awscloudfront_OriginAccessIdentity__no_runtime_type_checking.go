//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OriginAccessIdentity) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (o *jsiiProxy_OriginAccessIdentity) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (o *jsiiProxy_OriginAccessIdentity) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateOriginAccessIdentity_FromOriginAccessIdentityIdParameters(scope constructs.Construct, id *string, originAccessIdentityId *string) error {
	return nil
}

func validateOriginAccessIdentity_FromOriginAccessIdentityNameParameters(scope constructs.Construct, id *string, originAccessIdentityName *string) error {
	return nil
}

func validateOriginAccessIdentity_IsConstructParameters(x interface{}) error {
	return nil
}

func validateOriginAccessIdentity_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateOriginAccessIdentity_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewOriginAccessIdentityParameters(scope constructs.Construct, id *string, props *OriginAccessIdentityProps) error {
	return nil
}


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

func (o *jsiiProxy_OriginAccessIdentity) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (o *jsiiProxy_OriginAccessIdentity) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateOriginAccessIdentity_FromOriginAccessIdentityNameParameters(scope constructs.Construct, id *string, originAccessIdentityName *string) error {
	return nil
}

func validateOriginAccessIdentity_IsConstructParameters(x interface{}) error {
	return nil
}

func validateOriginAccessIdentity_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewOriginAccessIdentityParameters(scope constructs.Construct, id *string, props *OriginAccessIdentityProps) error {
	return nil
}


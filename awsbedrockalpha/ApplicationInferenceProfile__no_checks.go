//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationInferenceProfile) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_ApplicationInferenceProfile) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_ApplicationInferenceProfile) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationInferenceProfile) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_ApplicationInferenceProfile) validateGrantProfileUsageParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateApplicationInferenceProfile_FromApplicationInferenceProfileAttributesParameters(scope constructs.Construct, id *string, attrs *ApplicationInferenceProfileAttributes) error {
	return nil
}

func validateApplicationInferenceProfile_FromCfnApplicationInferenceProfileParameters(cfnApplicationInferenceProfile awsbedrock.CfnApplicationInferenceProfile) error {
	return nil
}

func validateApplicationInferenceProfile_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApplicationInferenceProfile_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateApplicationInferenceProfile_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewApplicationInferenceProfileParameters(scope constructs.Construct, id *string, props *ApplicationInferenceProfileProps) error {
	return nil
}


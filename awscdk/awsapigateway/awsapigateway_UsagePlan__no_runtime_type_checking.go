//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UsagePlan) validateAddApiKeyParameters(apiKey IApiKey, options *AddApiKeyOptions) error {
	return nil
}

func (u *jsiiProxy_UsagePlan) validateAddApiStageParameters(apiStage *UsagePlanPerApiStage) error {
	return nil
}

func (u *jsiiProxy_UsagePlan) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UsagePlan) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UsagePlan) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateUsagePlan_FromUsagePlanIdParameters(scope constructs.Construct, id *string, usagePlanId *string) error {
	return nil
}

func validateUsagePlan_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUsagePlan_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateUsagePlan_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewUsagePlanParameters(scope constructs.Construct, id *string, props *UsagePlanProps) error {
	return nil
}


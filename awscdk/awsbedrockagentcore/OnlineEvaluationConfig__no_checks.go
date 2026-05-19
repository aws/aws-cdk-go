//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OnlineEvaluationConfig) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (o *jsiiProxy_OnlineEvaluationConfig) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (o *jsiiProxy_OnlineEvaluationConfig) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (o *jsiiProxy_OnlineEvaluationConfig) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateOnlineEvaluationConfig_FromOnlineEvaluationConfigArnParameters(scope constructs.Construct, id *string, onlineEvaluationConfigArn *string) error {
	return nil
}

func validateOnlineEvaluationConfig_FromOnlineEvaluationConfigAttributesParameters(scope constructs.Construct, id *string, attrs *OnlineEvaluationConfigAttributes) error {
	return nil
}

func validateOnlineEvaluationConfig_FromOnlineEvaluationConfigIdParameters(scope constructs.Construct, id *string, onlineEvaluationConfigId *string) error {
	return nil
}

func validateOnlineEvaluationConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateOnlineEvaluationConfig_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateOnlineEvaluationConfig_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewOnlineEvaluationConfigParameters(scope constructs.Construct, id *string, props *OnlineEvaluationConfigProps) error {
	return nil
}


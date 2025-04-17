//go:build no_runtime_type_checking

package awsconfig

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) validateOnComplianceChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) validateOnReEvaluationStatusParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateCloudFormationStackDriftDetectionCheck_FromConfigRuleNameParameters(scope constructs.Construct, id *string, configRuleName *string) error {
	return nil
}

func validateCloudFormationStackDriftDetectionCheck_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCloudFormationStackDriftDetectionCheck_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCloudFormationStackDriftDetectionCheck_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCloudFormationStackDriftDetectionCheckParameters(scope constructs.Construct, id *string, props *CloudFormationStackDriftDetectionCheckProps) error {
	return nil
}


//go:build no_runtime_type_checking

package awsconfig

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFormationStackNotificationCheck) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CloudFormationStackNotificationCheck) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CloudFormationStackNotificationCheck) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_CloudFormationStackNotificationCheck) validateOnComplianceChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationStackNotificationCheck) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationStackNotificationCheck) validateOnReEvaluationStatusParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateCloudFormationStackNotificationCheck_FromConfigRuleNameParameters(scope constructs.Construct, id *string, configRuleName *string) error {
	return nil
}

func validateCloudFormationStackNotificationCheck_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCloudFormationStackNotificationCheck_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCloudFormationStackNotificationCheck_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCloudFormationStackNotificationCheckParameters(scope constructs.Construct, id *string, props *CloudFormationStackNotificationCheckProps) error {
	return nil
}


//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataQualityRuleset) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DataQualityRuleset) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DataQualityRuleset) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateDataQualityRuleset_FromRulesetArnParameters(scope constructs.Construct, id *string, rulesetArn *string) error {
	return nil
}

func validateDataQualityRuleset_FromRulesetNameParameters(scope constructs.Construct, id *string, rulesetName *string) error {
	return nil
}

func validateDataQualityRuleset_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDataQualityRuleset_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDataQualityRuleset_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDataQualityRulesetParameters(scope constructs.Construct, id *string, props *DataQualityRulesetProps) error {
	return nil
}


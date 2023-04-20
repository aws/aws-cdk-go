//go:build no_runtime_type_checking

package awsroute53resolver

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirewallRuleGroupAssociation) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateFirewallRuleGroupAssociation_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFirewallRuleGroupAssociation_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewFirewallRuleGroupAssociationParameters(scope constructs.Construct, id *string, props *FirewallRuleGroupAssociationProps) error {
	return nil
}


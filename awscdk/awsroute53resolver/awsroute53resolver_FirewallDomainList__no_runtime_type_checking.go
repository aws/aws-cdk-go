//go:build no_runtime_type_checking

package awsroute53resolver

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirewallDomainList) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FirewallDomainList) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FirewallDomainList) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (f *jsiiProxy_FirewallDomainList) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (f *jsiiProxy_FirewallDomainList) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateFirewallDomainList_FromFirewallDomainListIdParameters(scope constructs.Construct, id *string, firewallDomainListId *string) error {
	return nil
}

func validateFirewallDomainList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFirewallDomainList_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewFirewallDomainListParameters(scope constructs.Construct, id *string, props *FirewallDomainListProps) error {
	return nil
}


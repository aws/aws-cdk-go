//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityGroup) validateAddEgressRuleParameters(peer IPeer, connection Port) error {
	return nil
}

func (s *jsiiProxy_SecurityGroup) validateAddIngressRuleParameters(peer IPeer, connection Port) error {
	return nil
}

func (s *jsiiProxy_SecurityGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SecurityGroup) validateDetermineRuleScopeParameters(peer IPeer, connection Port, fromTo *string) error {
	return nil
}

func (s *jsiiProxy_SecurityGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SecurityGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSecurityGroup_FromLookupByIdParameters(scope constructs.Construct, id *string, securityGroupId *string) error {
	return nil
}

func validateSecurityGroup_FromLookupByNameParameters(scope constructs.Construct, id *string, securityGroupName *string, vpc IVpc) error {
	return nil
}

func validateSecurityGroup_FromSecurityGroupIdParameters(scope constructs.Construct, id *string, securityGroupId *string, options *SecurityGroupImportOptions) error {
	return nil
}

func validateSecurityGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecurityGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSecurityGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSecurityGroup_IsSecurityGroupParameters(x interface{}) error {
	return nil
}

func validateNewSecurityGroupParameters(scope constructs.Construct, id *string, props *SecurityGroupProps) error {
	return nil
}


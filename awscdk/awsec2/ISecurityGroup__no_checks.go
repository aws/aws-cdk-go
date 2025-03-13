//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ISecurityGroup) validateAddEgressRuleParameters(peer IPeer, connection Port) error {
	return nil
}

func (i *jsiiProxy_ISecurityGroup) validateAddIngressRuleParameters(peer IPeer, connection Port) error {
	return nil
}

func (i *jsiiProxy_ISecurityGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}


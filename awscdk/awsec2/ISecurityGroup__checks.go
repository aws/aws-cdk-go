//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (i *jsiiProxy_ISecurityGroup) validateAddEgressRuleParameters(peer IPeer, connection Port) error {
	if peer == nil {
		return fmt.Errorf("parameter peer is required, but nil was provided")
	}

	if connection == nil {
		return fmt.Errorf("parameter connection is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ISecurityGroup) validateAddIngressRuleParameters(peer IPeer, connection Port) error {
	if peer == nil {
		return fmt.Errorf("parameter peer is required, but nil was provided")
	}

	if connection == nil {
		return fmt.Errorf("parameter connection is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ISecurityGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}


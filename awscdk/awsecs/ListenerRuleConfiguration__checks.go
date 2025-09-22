//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

func validateListenerRuleConfiguration_ApplicationListenerRuleParameters(rule awselasticloadbalancingv2.ApplicationListenerRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}

	return nil
}

func validateListenerRuleConfiguration_NetworkListenerParameters(listener awselasticloadbalancingv2.NetworkListener) error {
	if listener == nil {
		return fmt.Errorf("parameter listener is required, but nil was provided")
	}

	return nil
}


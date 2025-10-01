//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func validateListenerRuleConfiguration_ApplicationListenerRuleParameters(rule awselasticloadbalancingv2.ApplicationListenerRule) error {
	return nil
}

func validateListenerRuleConfiguration_NetworkListenerParameters(listener awselasticloadbalancingv2.NetworkListener) error {
	return nil
}


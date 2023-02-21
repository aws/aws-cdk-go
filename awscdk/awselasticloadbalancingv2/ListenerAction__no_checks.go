//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_ListenerAction) validateBindParameters(scope constructs.Construct, listener IApplicationListener) error {
	return nil
}

func (l *jsiiProxy_ListenerAction) validateRenumberParameters(actions *[]*CfnListener_ActionProperty) error {
	return nil
}

func validateListenerAction_AuthenticateOidcParameters(options *AuthenticateOidcOptions) error {
	return nil
}

func validateListenerAction_FixedResponseParameters(statusCode *float64, options *FixedResponseOptions) error {
	return nil
}

func validateListenerAction_ForwardParameters(targetGroups *[]IApplicationTargetGroup, options *ForwardOptions) error {
	return nil
}

func validateListenerAction_RedirectParameters(options *RedirectOptions) error {
	return nil
}

func validateListenerAction_WeightedForwardParameters(targetGroups *[]*WeightedTargetGroup, options *ForwardOptions) error {
	return nil
}

func validateNewListenerActionParameters(actionJson *CfnListener_ActionProperty) error {
	return nil
}


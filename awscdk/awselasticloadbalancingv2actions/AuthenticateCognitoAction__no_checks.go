//go:build no_runtime_type_checking

package awselasticloadbalancingv2actions

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AuthenticateCognitoAction) validateBindParameters(scope constructs.Construct, listener awselasticloadbalancingv2.IApplicationListener) error {
	return nil
}

func (a *jsiiProxy_AuthenticateCognitoAction) validateRenumberParameters(actions *[]*awselasticloadbalancingv2.CfnListener_ActionProperty) error {
	return nil
}

func validateAuthenticateCognitoAction_AuthenticateOidcParameters(options *awselasticloadbalancingv2.AuthenticateOidcOptions) error {
	return nil
}

func validateAuthenticateCognitoAction_FixedResponseParameters(statusCode *float64, options *awselasticloadbalancingv2.FixedResponseOptions) error {
	return nil
}

func validateAuthenticateCognitoAction_ForwardParameters(targetGroups *[]awselasticloadbalancingv2.IApplicationTargetGroup, options *awselasticloadbalancingv2.ForwardOptions) error {
	return nil
}

func validateAuthenticateCognitoAction_RedirectParameters(options *awselasticloadbalancingv2.RedirectOptions) error {
	return nil
}

func validateAuthenticateCognitoAction_WeightedForwardParameters(targetGroups *[]*awselasticloadbalancingv2.WeightedTargetGroup, options *awselasticloadbalancingv2.ForwardOptions) error {
	return nil
}

func validateNewAuthenticateCognitoActionParameters(options *AuthenticateCognitoActionProps) error {
	return nil
}


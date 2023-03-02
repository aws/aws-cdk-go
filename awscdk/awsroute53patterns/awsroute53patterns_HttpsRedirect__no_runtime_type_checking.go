//go:build no_runtime_type_checking

package awsroute53patterns

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpsRedirect) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (h *jsiiProxy_HttpsRedirect) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateHttpsRedirect_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewHttpsRedirectParameters(scope constructs.Construct, id *string, props *HttpsRedirectProps) error {
	return nil
}


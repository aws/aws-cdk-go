//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Resolver) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (r *jsiiProxy_Resolver) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateResolver_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewResolverParameters(scope constructs.Construct, id *string, props *ResolverProps) error {
	return nil
}


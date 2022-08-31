//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BaseJenkinsProvider) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (b *jsiiProxy_BaseJenkinsProvider) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateBaseJenkinsProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewBaseJenkinsProviderParameters(scope constructs.Construct, id *string) error {
	return nil
}


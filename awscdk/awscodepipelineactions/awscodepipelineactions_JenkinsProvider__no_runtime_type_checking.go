//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JenkinsProvider) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (j *jsiiProxy_JenkinsProvider) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateJenkinsProvider_FromJenkinsProviderAttributesParameters(scope constructs.Construct, id *string, attrs *JenkinsProviderAttributes) error {
	return nil
}

func validateJenkinsProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewJenkinsProviderParameters(scope constructs.Construct, id *string, props *JenkinsProviderProps) error {
	return nil
}


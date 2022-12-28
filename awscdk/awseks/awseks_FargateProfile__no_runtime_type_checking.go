//go:build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FargateProfile) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (f *jsiiProxy_FargateProfile) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateFargateProfile_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewFargateProfileParameters(scope constructs.Construct, id *string, props *FargateProfileProps) error {
	return nil
}


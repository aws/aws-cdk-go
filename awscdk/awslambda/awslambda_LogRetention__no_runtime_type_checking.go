//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogRetention) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (l *jsiiProxy_LogRetention) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateLogRetention_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLogRetentionParameters(scope constructs.Construct, id *string, props *LogRetentionProps) error {
	return nil
}


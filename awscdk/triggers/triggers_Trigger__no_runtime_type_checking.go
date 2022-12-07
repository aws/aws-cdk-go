//go:build no_runtime_type_checking

package triggers

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Trigger) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (t *jsiiProxy_Trigger) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateTrigger_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTriggerParameters(scope constructs.Construct, id *string, props *TriggerProps) error {
	return nil
}


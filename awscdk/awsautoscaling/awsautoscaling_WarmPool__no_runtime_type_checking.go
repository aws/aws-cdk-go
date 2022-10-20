//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WarmPool) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WarmPool) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WarmPool) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (w *jsiiProxy_WarmPool) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (w *jsiiProxy_WarmPool) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateWarmPool_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWarmPool_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewWarmPoolParameters(scope constructs.Construct, id *string, props *WarmPoolProps) error {
	return nil
}


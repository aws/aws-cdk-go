//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HelmChart) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (h *jsiiProxy_HelmChart) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateHelmChart_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewHelmChartParameters(scope constructs.Construct, id *string, props *HelmChartProps) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awssynthetics

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Canary) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_Canary) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_Canary) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_Canary) validateMetricDurationParameters(options *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_Canary) validateMetricFailedParameters(options *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_Canary) validateMetricSuccessPercentParameters(options *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_Canary) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_Canary) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateCanary_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCanary_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewCanaryParameters(scope constructs.Construct, id *string, props *CanaryProps) error {
	return nil
}


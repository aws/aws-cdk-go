//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicyEngineBase) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (p *jsiiProxy_PolicyEngineBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PolicyEngineBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PolicyEngineBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_PolicyEngineBase) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (p *jsiiProxy_PolicyEngineBase) validateGrantEvaluateParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (p *jsiiProxy_PolicyEngineBase) validateGrantEvaluateForGatewayParameters(grantee awsiam.IGrantable, gateway IGateway) error {
	return nil
}

func (p *jsiiProxy_PolicyEngineBase) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (p *jsiiProxy_PolicyEngineBase) validateMetricParameters(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_PolicyEngineBase) validateMetricAuthorizationLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_PolicyEngineBase) validateMetricAuthorizationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_PolicyEngineBase) validateMetricDeniedRequestsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_PolicyEngineBase) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validatePolicyEngineBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePolicyEngineBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePolicyEngineBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPolicyEngineBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}


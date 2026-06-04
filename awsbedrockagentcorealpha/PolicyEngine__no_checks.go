//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicyEngine) validateAddPolicyParameters(id *string, options *AddPolicyOptions) error {
	return nil
}

func (p *jsiiProxy_PolicyEngine) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (p *jsiiProxy_PolicyEngine) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PolicyEngine) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PolicyEngine) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_PolicyEngine) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (p *jsiiProxy_PolicyEngine) validateGrantEvaluateParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (p *jsiiProxy_PolicyEngine) validateGrantEvaluateForGatewayParameters(grantee awsiam.IGrantable, gateway IGateway) error {
	return nil
}

func (p *jsiiProxy_PolicyEngine) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (p *jsiiProxy_PolicyEngine) validateMetricParameters(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_PolicyEngine) validateMetricAuthorizationLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_PolicyEngine) validateMetricAuthorizationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_PolicyEngine) validateMetricDeniedRequestsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_PolicyEngine) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validatePolicyEngine_FromPolicyEngineAttributesParameters(scope constructs.Construct, id *string, attrs *PolicyEngineAttributes) error {
	return nil
}

func validatePolicyEngine_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePolicyEngine_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePolicyEngine_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPolicyEngineParameters(scope constructs.Construct, id *string, props *PolicyEngineProps) error {
	return nil
}


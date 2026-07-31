//go:build no_runtime_type_checking

package awsmediaconnectalpha

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_Gateway) validateAddNetworkParameters(network GatewayNetwork) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateMetricEgressBridgeDroppedPacketsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateMetricEgressBridgeTotalPacketsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateMetricIngressBridgeDroppedPacketsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateMetricIngressBridgeTotalPacketsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateGateway_FromGatewayArnParameters(scope constructs.Construct, id *string, gatewayArn *string) error {
	return nil
}

func validateGateway_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGateway_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGateway_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGatewayParameters(scope constructs.Construct, id *string, props *GatewayProps) error {
	return nil
}


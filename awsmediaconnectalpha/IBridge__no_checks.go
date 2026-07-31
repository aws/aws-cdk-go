//go:build no_runtime_type_checking

package awsmediaconnectalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IBridge) validateAddOutputParameters(id *string, networkOutput *BridgeNetworkOutput) error {
	return nil
}

func (i *jsiiProxy_IBridge) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IBridge) validateMetricFailoverSwitchesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IBridge) validateMetricSourceBitrateParameters(bridgeSourceName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IBridge) validateMetricSourcePacketLossPercentParameters(bridgeSourceName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IBridge) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}


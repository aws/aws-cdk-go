//go:build no_runtime_type_checking

package awsmediaconnectalpha

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_Bridge) validateAddOutputParameters(id *string, networkOutput *BridgeNetworkOutput) error {
	return nil
}

func (b *jsiiProxy_Bridge) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (b *jsiiProxy_Bridge) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_Bridge) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_Bridge) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (b *jsiiProxy_Bridge) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_Bridge) validateMetricFailoverSwitchesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_Bridge) validateMetricSourceBitrateParameters(bridgeSourceName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_Bridge) validateMetricSourcePacketLossPercentParameters(bridgeSourceName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateBridge_FromBridgeArnParameters(scope constructs.Construct, id *string, bridgeArn *string) error {
	return nil
}

func validateBridge_FromBridgeAttributesParameters(scope constructs.Construct, id *string, attrs *BridgeAttributes) error {
	return nil
}

func validateBridge_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBridge_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBridge_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBridgeParameters(scope constructs.Construct, id *string, props *BridgeProps) error {
	return nil
}


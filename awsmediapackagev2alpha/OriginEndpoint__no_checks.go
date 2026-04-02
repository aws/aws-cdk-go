//go:build no_runtime_type_checking

package awsmediapackagev2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OriginEndpoint) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement, cdnAuth *CdnAuthConfiguration) error {
	return nil
}

func (o *jsiiProxy_OriginEndpoint) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (o *jsiiProxy_OriginEndpoint) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (o *jsiiProxy_OriginEndpoint) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (o *jsiiProxy_OriginEndpoint) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (o *jsiiProxy_OriginEndpoint) validateMetricEgressBytesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (o *jsiiProxy_OriginEndpoint) validateMetricEgressRequestCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (o *jsiiProxy_OriginEndpoint) validateMetricEgressResponseTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (o *jsiiProxy_OriginEndpoint) validateMetricIngressBytesParameters(options *awscloudwatch.MetricOptions) error {
	return nil
}

func (o *jsiiProxy_OriginEndpoint) validateMetricIngressRequestCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (o *jsiiProxy_OriginEndpoint) validateMetricIngressResponseTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (o *jsiiProxy_OriginEndpoint) validateSegmentValidationParameters(segmentContainerType ContainerType, segment *SegmentConfiguration) error {
	return nil
}

func validateOriginEndpoint_FromOriginEndpointAttributesParameters(scope constructs.Construct, id *string, attrs *OriginEndpointAttributes) error {
	return nil
}

func validateOriginEndpoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateOriginEndpoint_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateOriginEndpoint_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_OriginEndpoint) validateSetAutoCreatePolicyParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_OriginEndpoint) validateSetDashManifestsParameters(val *[]*awsmediapackagev2.CfnOriginEndpoint_DashManifestConfigurationProperty) error {
	return nil
}

func (j *jsiiProxy_OriginEndpoint) validateSetHlsManifestsParameters(val *[]*awsmediapackagev2.CfnOriginEndpoint_HlsManifestConfigurationProperty) error {
	return nil
}

func (j *jsiiProxy_OriginEndpoint) validateSetLlHlsManifestsParameters(val *[]*awsmediapackagev2.CfnOriginEndpoint_LowLatencyHlsManifestConfigurationProperty) error {
	return nil
}

func (j *jsiiProxy_OriginEndpoint) validateSetMssManifestsParameters(val *[]*awsmediapackagev2.CfnOriginEndpoint_MssManifestConfigurationProperty) error {
	return nil
}

func (j *jsiiProxy_OriginEndpoint) validateSetSegmentParameters(val *SegmentConfiguration) error {
	return nil
}

func validateNewOriginEndpointParameters(scope constructs.Construct, id *string, props *OriginEndpointProps) error {
	return nil
}


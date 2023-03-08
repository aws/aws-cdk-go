//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalService) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateAssociateCloudMapServiceParameters(_options *AssociateCloudMapServiceOptions) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateAttachToApplicationTargetGroupParameters(_targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateAttachToClassicLBParameters(loadBalancer awselasticloadbalancing.LoadBalancer) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateAutoScaleTaskCountParameters(_props *awsapplicationautoscaling.EnableScalingProps) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateConfigureAwsVpcNetworkingWithSecurityGroupsParameters(_vpc awsec2.IVpc, _vpcSubnets *awsec2.SubnetSelection) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateEnableCloudMapParameters(_options *CloudMapOptions) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateEnableServiceConnectParameters(config *ServiceConnectProps) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateLoadBalancerTargetParameters(_options *LoadBalancerTargetOptions) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateMetricCpuUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateMetricMemoryUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateRegisterLoadBalancerTargetsParameters(_targets *[]*EcsTarget) error {
	return nil
}

func validateExternalService_FromExternalServiceArnParameters(scope constructs.Construct, id *string, externalServiceArn *string) error {
	return nil
}

func validateExternalService_FromExternalServiceAttributesParameters(scope constructs.Construct, id *string, attrs *ExternalServiceAttributes) error {
	return nil
}

func validateExternalService_FromServiceArnWithClusterParameters(scope constructs.Construct, id *string, serviceArn *string) error {
	return nil
}

func validateExternalService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateExternalService_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateExternalService_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_ExternalService) validateSetLoadBalancersParameters(val *[]*CfnService_LoadBalancerProperty) error {
	return nil
}

func (j *jsiiProxy_ExternalService) validateSetNetworkConfigurationParameters(val *CfnService_NetworkConfigurationProperty) error {
	return nil
}

func (j *jsiiProxy_ExternalService) validateSetServiceRegistriesParameters(val *[]*CfnService_ServiceRegistryProperty) error {
	return nil
}

func validateNewExternalServiceParameters(scope constructs.Construct, id *string, props *ExternalServiceProps) error {
	return nil
}


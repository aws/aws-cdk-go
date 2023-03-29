//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FargateService) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateAssociateCloudMapServiceParameters(options *AssociateCloudMapServiceOptions) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateAttachToApplicationTargetGroupParameters(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateAttachToClassicLBParameters(loadBalancer awselasticloadbalancing.LoadBalancer) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateAutoScaleTaskCountParameters(props *awsapplicationautoscaling.EnableScalingProps) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateConfigureAwsVpcNetworkingWithSecurityGroupsParameters(vpc awsec2.IVpc, vpcSubnets *awsec2.SubnetSelection) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateEnableCloudMapParameters(options *CloudMapOptions) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateEnableServiceConnectParameters(config *ServiceConnectProps) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateLoadBalancerTargetParameters(options *LoadBalancerTargetOptions) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateMetricCpuUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateMetricMemoryUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (f *jsiiProxy_FargateService) validateRegisterLoadBalancerTargetsParameters(targets *[]*EcsTarget) error {
	return nil
}

func validateFargateService_FromFargateServiceArnParameters(scope constructs.Construct, id *string, fargateServiceArn *string) error {
	return nil
}

func validateFargateService_FromFargateServiceAttributesParameters(scope constructs.Construct, id *string, attrs *FargateServiceAttributes) error {
	return nil
}

func validateFargateService_FromServiceArnWithClusterParameters(scope constructs.Construct, id *string, serviceArn *string) error {
	return nil
}

func validateFargateService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFargateService_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFargateService_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_FargateService) validateSetLoadBalancersParameters(val *[]*CfnService_LoadBalancerProperty) error {
	return nil
}

func (j *jsiiProxy_FargateService) validateSetNetworkConfigurationParameters(val *CfnService_NetworkConfigurationProperty) error {
	return nil
}

func (j *jsiiProxy_FargateService) validateSetServiceRegistriesParameters(val *[]*CfnService_ServiceRegistryProperty) error {
	return nil
}

func validateNewFargateServiceParameters(scope constructs.Construct, id *string, props *FargateServiceProps) error {
	return nil
}


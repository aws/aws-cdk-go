//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BaseService) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateAssociateCloudMapServiceParameters(options *AssociateCloudMapServiceOptions) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateAttachToApplicationTargetGroupParameters(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateAttachToClassicLBParameters(loadBalancer awselasticloadbalancing.LoadBalancer) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateAutoScaleTaskCountParameters(props *awsapplicationautoscaling.EnableScalingProps) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateConfigureAwsVpcNetworkingWithSecurityGroupsParameters(vpc awsec2.IVpc, vpcSubnets *awsec2.SubnetSelection) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateEnableCloudMapParameters(options *CloudMapOptions) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateEnableServiceConnectParameters(config *ServiceConnectProps) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateLoadBalancerTargetParameters(options *LoadBalancerTargetOptions) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateMetricCpuUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateMetricMemoryUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BaseService) validateRegisterLoadBalancerTargetsParameters(targets *[]*EcsTarget) error {
	return nil
}

func validateBaseService_FromServiceArnWithClusterParameters(scope constructs.Construct, id *string, serviceArn *string) error {
	return nil
}

func validateBaseService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBaseService_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBaseService_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_BaseService) validateSetLoadBalancersParameters(val *[]*CfnService_LoadBalancerProperty) error {
	return nil
}

func (j *jsiiProxy_BaseService) validateSetNetworkConfigurationParameters(val *CfnService_NetworkConfigurationProperty) error {
	return nil
}

func (j *jsiiProxy_BaseService) validateSetServiceRegistriesParameters(val *[]*CfnService_ServiceRegistryProperty) error {
	return nil
}

func validateNewBaseServiceParameters(scope constructs.Construct, id *string, props *BaseServiceProps, additionalProps interface{}, taskDefinition TaskDefinition) error {
	return nil
}


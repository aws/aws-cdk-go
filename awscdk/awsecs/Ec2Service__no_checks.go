//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2Service) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateAssociateCloudMapServiceParameters(options *AssociateCloudMapServiceOptions) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateAttachToApplicationTargetGroupParameters(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateAttachToClassicLBParameters(loadBalancer awselasticloadbalancing.LoadBalancer) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateAutoScaleTaskCountParameters(props *awsapplicationautoscaling.EnableScalingProps) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateConfigureAwsVpcNetworkingWithSecurityGroupsParameters(vpc awsec2.IVpc, vpcSubnets *awsec2.SubnetSelection) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateEnableCloudMapParameters(options *CloudMapOptions) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateEnableServiceConnectParameters(config *ServiceConnectProps) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateLoadBalancerTargetParameters(options *LoadBalancerTargetOptions) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateMetricCpuUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateMetricMemoryUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_Ec2Service) validateRegisterLoadBalancerTargetsParameters(targets *[]*EcsTarget) error {
	return nil
}

func validateEc2Service_FromEc2ServiceArnParameters(scope constructs.Construct, id *string, ec2ServiceArn *string) error {
	return nil
}

func validateEc2Service_FromEc2ServiceAttributesParameters(scope constructs.Construct, id *string, attrs *Ec2ServiceAttributes) error {
	return nil
}

func validateEc2Service_FromServiceArnWithClusterParameters(scope constructs.Construct, id *string, serviceArn *string) error {
	return nil
}

func validateEc2Service_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEc2Service_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEc2Service_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_Ec2Service) validateSetLoadBalancersParameters(val *[]*CfnService_LoadBalancerProperty) error {
	return nil
}

func (j *jsiiProxy_Ec2Service) validateSetNetworkConfigurationParameters(val *CfnService_NetworkConfigurationProperty) error {
	return nil
}

func (j *jsiiProxy_Ec2Service) validateSetServiceRegistriesParameters(val *[]*CfnService_ServiceRegistryProperty) error {
	return nil
}

func validateNewEc2ServiceParameters(scope constructs.Construct, id *string, props *Ec2ServiceProps) error {
	return nil
}


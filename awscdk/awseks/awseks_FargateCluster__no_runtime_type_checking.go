//go:build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FargateCluster) validateAddAutoScalingGroupCapacityParameters(id *string, options *AutoScalingGroupCapacityOptions) error {
	return nil
}

func (f *jsiiProxy_FargateCluster) validateAddCdk8sChartParameters(id *string, chart constructs.Construct, options *KubernetesManifestOptions) error {
	return nil
}

func (f *jsiiProxy_FargateCluster) validateAddFargateProfileParameters(id *string, options *FargateProfileOptions) error {
	return nil
}

func (f *jsiiProxy_FargateCluster) validateAddHelmChartParameters(id *string, options *HelmChartOptions) error {
	return nil
}

func (f *jsiiProxy_FargateCluster) validateAddManifestParameters(id *string) error {
	return nil
}

func (f *jsiiProxy_FargateCluster) validateAddNodegroupCapacityParameters(id *string, options *NodegroupOptions) error {
	return nil
}

func (f *jsiiProxy_FargateCluster) validateAddServiceAccountParameters(id *string, options *ServiceAccountOptions) error {
	return nil
}

func (f *jsiiProxy_FargateCluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FargateCluster) validateConnectAutoScalingGroupCapacityParameters(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions) error {
	return nil
}

func (f *jsiiProxy_FargateCluster) validateGetIngressLoadBalancerAddressParameters(ingressName *string, options *IngressLoadBalancerAddressOptions) error {
	return nil
}

func (f *jsiiProxy_FargateCluster) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FargateCluster) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (f *jsiiProxy_FargateCluster) validateGetServiceLoadBalancerAddressParameters(serviceName *string, options *ServiceLoadBalancerAddressOptions) error {
	return nil
}

func validateFargateCluster_FromClusterAttributesParameters(scope constructs.Construct, id *string, attrs *ClusterAttributes) error {
	return nil
}

func validateFargateCluster_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFargateCluster_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFargateCluster_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFargateClusterParameters(scope constructs.Construct, id *string, props *FargateClusterProps) error {
	return nil
}


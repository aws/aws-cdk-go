//go:build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Cluster) validateAddAutoScalingGroupCapacityParameters(id *string, options *AutoScalingGroupCapacityOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddCdk8sChartParameters(id *string, chart constructs.Construct, options *KubernetesManifestOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddFargateProfileParameters(id *string, options *FargateProfileOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddHelmChartParameters(id *string, options *HelmChartOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddManifestParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddNodegroupCapacityParameters(id *string, options *NodegroupOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddServiceAccountParameters(id *string, options *ServiceAccountOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateConnectAutoScalingGroupCapacityParameters(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetIngressLoadBalancerAddressParameters(ingressName *string, options *IngressLoadBalancerAddressOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetServiceLoadBalancerAddressParameters(serviceName *string, options *ServiceLoadBalancerAddressOptions) error {
	return nil
}

func validateCluster_FromClusterAttributesParameters(scope constructs.Construct, id *string, attrs *ClusterAttributes) error {
	return nil
}

func validateCluster_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCluster_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCluster_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewClusterParameters(scope constructs.Construct, id *string, props *ClusterProps) error {
	return nil
}


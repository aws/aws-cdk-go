//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Cluster) validateAddAsgCapacityProviderParameters(provider AsgCapacityProvider, options *AddAutoScalingGroupCapacityOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddAutoScalingGroupParameters(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AddAutoScalingGroupCapacityOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddCapacityParameters(id *string, options *AddCapacityOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddCapacityProviderParameters(provider *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddDefaultCloudMapNamespaceParameters(options *CloudMapNamespaceOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateMetricCpuReservationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateMetricCpuUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateMetricMemoryReservationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateMetricMemoryUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateCluster_FromClusterArnParameters(scope constructs.Construct, id *string, clusterArn *string) error {
	return nil
}

func validateCluster_FromClusterAttributesParameters(scope constructs.Construct, id *string, attrs *ClusterAttributes) error {
	return nil
}

func validateCluster_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCluster_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewClusterParameters(scope constructs.Construct, id *string, props *ClusterProps) error {
	return nil
}


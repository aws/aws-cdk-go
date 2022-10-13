//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsekslegacy

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Cluster) validateAddAutoScalingGroupParameters(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddCapacityParameters(id *string, options *CapacityOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddChartParameters(id *string, options *HelmChartOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddResourceParameters(id *string) error {
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

func (c *jsiiProxy_Cluster) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateCluster_FromClusterAttributesParameters(scope awscdk.Construct, id *string, attrs *ClusterAttributes) error {
	return nil
}

func validateCluster_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCluster_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewClusterParameters(scope awscdk.Construct, id *string, props *ClusterProps) error {
	return nil
}


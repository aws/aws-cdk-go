//go:build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ICluster) validateAddCdk8sChartParameters(id *string, chart constructs.Construct, options *KubernetesManifestOptions) error {
	return nil
}

func (i *jsiiProxy_ICluster) validateAddHelmChartParameters(id *string, options *HelmChartOptions) error {
	return nil
}

func (i *jsiiProxy_ICluster) validateAddManifestParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_ICluster) validateAddServiceAccountParameters(id *string, options *ServiceAccountOptions) error {
	return nil
}

func (i *jsiiProxy_ICluster) validateConnectAutoScalingGroupCapacityParameters(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions) error {
	return nil
}

func (i *jsiiProxy_ICluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}


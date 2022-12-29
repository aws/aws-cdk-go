//go:build no_runtime_type_checking

package awsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutoScalingGroup) validateAddLifecycleHookParameters(id *string, props *BasicLifecycleHookProps) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateAddSecurityGroupParameters(securityGroup awsec2.ISecurityGroup) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateAddWarmPoolParameters(options *WarmPoolOptions) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateApplyCloudFormationInitParameters(init awsec2.CloudFormationInit, options *ApplyCloudFormationInitOptions) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateAttachToApplicationTargetGroupParameters(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateAttachToClassicLBParameters(loadBalancer awselasticloadbalancing.LoadBalancer) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateScaleOnCpuUtilizationParameters(id *string, props *CpuUtilizationScalingProps) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateScaleOnIncomingBytesParameters(id *string, props *NetworkUtilizationScalingProps) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateScaleOnMetricParameters(id *string, props *BasicStepScalingPolicyProps) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateScaleOnOutgoingBytesParameters(id *string, props *NetworkUtilizationScalingProps) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateScaleOnRequestCountParameters(id *string, props *RequestCountScalingProps) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateScaleOnScheduleParameters(id *string, props *BasicScheduledActionProps) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroup) validateScaleToTrackMetricParameters(id *string, props *MetricTargetTrackingProps) error {
	return nil
}

func validateAutoScalingGroup_FromAutoScalingGroupNameParameters(scope constructs.Construct, id *string, autoScalingGroupName *string) error {
	return nil
}

func validateAutoScalingGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAutoScalingGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAutoScalingGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_AutoScalingGroup) validateSetHasCalledScaleOnRequestCountParameters(val *bool) error {
	return nil
}

func validateNewAutoScalingGroupParameters(scope constructs.Construct, id *string, props *AutoScalingGroupProps) error {
	return nil
}


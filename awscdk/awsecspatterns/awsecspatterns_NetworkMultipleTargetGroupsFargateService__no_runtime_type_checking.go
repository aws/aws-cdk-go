//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) validateAddPortMappingForTargetsParameters(container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) error {
	return nil
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) validateRegisterECSTargetsParameters(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) error {
	return nil
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateNetworkMultipleTargetGroupsFargateService_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) validateSetListenersParameters(val *[]awselasticloadbalancingv2.NetworkListener) error {
	return nil
}

func (j *jsiiProxy_NetworkMultipleTargetGroupsFargateService) validateSetTargetGroupsParameters(val *[]awselasticloadbalancingv2.NetworkTargetGroup) error {
	return nil
}

func validateNewNetworkMultipleTargetGroupsFargateServiceParameters(scope constructs.Construct, id *string, props *NetworkMultipleTargetGroupsFargateServiceProps) error {
	return nil
}


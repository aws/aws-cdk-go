//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) validateAddPortMappingForTargetsParameters(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) error {
	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) validateRegisterECSTargetsParameters(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) error {
	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateApplicationMultipleTargetGroupsEc2Service_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) validateSetListenersParameters(val *[]awselasticloadbalancingv2.ApplicationListener) error {
	return nil
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsEc2Service) validateSetTargetGroupsParameters(val *[]awselasticloadbalancingv2.ApplicationTargetGroup) error {
	return nil
}

func validateNewApplicationMultipleTargetGroupsEc2ServiceParameters(scope constructs.Construct, id *string, props *ApplicationMultipleTargetGroupsEc2ServiceProps) error {
	return nil
}


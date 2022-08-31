//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) validateAddPortMappingForTargetsParameters(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) error {
	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) validateRegisterECSTargetsParameters(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) error {
	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateApplicationMultipleTargetGroupsServiceBase_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) validateSetListenersParameters(val *[]awselasticloadbalancingv2.ApplicationListener) error {
	return nil
}

func (j *jsiiProxy_ApplicationMultipleTargetGroupsServiceBase) validateSetTargetGroupsParameters(val *[]awselasticloadbalancingv2.ApplicationTargetGroup) error {
	return nil
}

func validateNewApplicationMultipleTargetGroupsServiceBaseParameters(scope constructs.Construct, id *string, props *ApplicationMultipleTargetGroupsServiceBaseProps) error {
	return nil
}


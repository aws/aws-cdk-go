//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) validateAddPortMappingForTargetsParameters(container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) error {
	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func (a *jsiiProxy_ApplicationMultipleTargetGroupsFargateService) validateRegisterECSTargetsParameters(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*ApplicationTargetProps) error {
	return nil
}

func validateApplicationMultipleTargetGroupsFargateService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewApplicationMultipleTargetGroupsFargateServiceParameters(scope constructs.Construct, id *string, props *ApplicationMultipleTargetGroupsFargateServiceProps) error {
	return nil
}


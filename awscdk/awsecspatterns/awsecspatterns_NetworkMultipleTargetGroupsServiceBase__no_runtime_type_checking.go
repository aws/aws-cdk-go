//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) validateAddPortMappingForTargetsParameters(container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) error {
	return nil
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsServiceBase) validateRegisterECSTargetsParameters(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) error {
	return nil
}

func validateNetworkMultipleTargetGroupsServiceBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewNetworkMultipleTargetGroupsServiceBaseParameters(scope constructs.Construct, id *string, props *NetworkMultipleTargetGroupsServiceBaseProps) error {
	return nil
}


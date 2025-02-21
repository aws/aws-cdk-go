//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) validateAddPortMappingForTargetsParameters(container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) error {
	return nil
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsEc2Service) validateRegisterECSTargetsParameters(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) error {
	return nil
}

func validateNetworkMultipleTargetGroupsEc2Service_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewNetworkMultipleTargetGroupsEc2ServiceParameters(scope constructs.Construct, id *string, props *NetworkMultipleTargetGroupsEc2ServiceProps) error {
	return nil
}


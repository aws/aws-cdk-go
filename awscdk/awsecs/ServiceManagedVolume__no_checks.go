//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceManagedVolume) validateMountInParameters(container ContainerDefinition, mountPoint *ContainerMountPoint) error {
	return nil
}

func validateServiceManagedVolume_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewServiceManagedVolumeParameters(scope constructs.Construct, id *string, props *ServiceManagedVolumeProps) error {
	return nil
}


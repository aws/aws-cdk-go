//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirelensLogRouter) validateAddContainerDependenciesParameters(containerDependencies *[]*ContainerDependency) error {
	return nil
}

func (f *jsiiProxy_FirelensLogRouter) validateAddEnvironmentParameters(name *string, value *string) error {
	return nil
}

func (f *jsiiProxy_FirelensLogRouter) validateAddLinkParameters(container ContainerDefinition) error {
	return nil
}

func (f *jsiiProxy_FirelensLogRouter) validateAddMountPointsParameters(mountPoints *[]*MountPoint) error {
	return nil
}

func (f *jsiiProxy_FirelensLogRouter) validateAddPortMappingsParameters(portMappings *[]*PortMapping) error {
	return nil
}

func (f *jsiiProxy_FirelensLogRouter) validateAddScratchParameters(scratch *ScratchSpace) error {
	return nil
}

func (f *jsiiProxy_FirelensLogRouter) validateAddSecretParameters(name *string, secret Secret) error {
	return nil
}

func (f *jsiiProxy_FirelensLogRouter) validateAddToExecutionPolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (f *jsiiProxy_FirelensLogRouter) validateAddUlimitsParameters(ulimits *[]*Ulimit) error {
	return nil
}

func (f *jsiiProxy_FirelensLogRouter) validateAddVolumesFromParameters(volumesFrom *[]*VolumeFrom) error {
	return nil
}

func (f *jsiiProxy_FirelensLogRouter) validateFindPortMappingParameters(containerPort *float64, protocol Protocol) error {
	return nil
}

func (f *jsiiProxy_FirelensLogRouter) validateFindPortMappingByNameParameters(name *string) error {
	return nil
}

func validateFirelensLogRouter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewFirelensLogRouterParameters(scope constructs.Construct, id *string, props *FirelensLogRouterProps) error {
	return nil
}


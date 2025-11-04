//go:build no_runtime_type_checking

package awscdkapplicationsignalsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DotNetLinuxInjector) validateInjectAdditionalEnvironmentsParameters(envsToInject *map[string]*string, envsFromTaskDef *map[string]*string) error {
	return nil
}

func (d *jsiiProxy_DotNetLinuxInjector) validateInjectInitContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (d *jsiiProxy_DotNetLinuxInjector) validateOverrideAdditionalEnvironmentsParameters(envsToOverride *map[string]*string, envsFromTaskDef *map[string]*string) error {
	return nil
}

func (d *jsiiProxy_DotNetLinuxInjector) validateRenderDefaultContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (j *jsiiProxy_DotNetLinuxInjector) validateSetInstrumentationVersionParameters(val InstrumentationVersion) error {
	return nil
}

func (j *jsiiProxy_DotNetLinuxInjector) validateSetSharedVolumeNameParameters(val *string) error {
	return nil
}

func validateNewDotNetLinuxInjectorParameters(sharedVolumeName *string, instrumentationVersion InstrumentationVersion, cpuArch awsecs.CpuArchitecture, overrideEnvironments *[]*EnvironmentExtension) error {
	return nil
}


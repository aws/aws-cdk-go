//go:build no_runtime_type_checking

package awscdkapplicationsignalsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DotNetWindowsInjector) validateInjectAdditionalEnvironmentsParameters(envsToInject *map[string]*string, envsFromTaskDef *map[string]*string) error {
	return nil
}

func (d *jsiiProxy_DotNetWindowsInjector) validateInjectInitContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (d *jsiiProxy_DotNetWindowsInjector) validateOverrideAdditionalEnvironmentsParameters(envsToOverride *map[string]*string, envsFromTaskDef *map[string]*string) error {
	return nil
}

func (d *jsiiProxy_DotNetWindowsInjector) validateRenderDefaultContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (j *jsiiProxy_DotNetWindowsInjector) validateSetInstrumentationVersionParameters(val InstrumentationVersion) error {
	return nil
}

func (j *jsiiProxy_DotNetWindowsInjector) validateSetSharedVolumeNameParameters(val *string) error {
	return nil
}

func validateNewDotNetWindowsInjectorParameters(sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) error {
	return nil
}


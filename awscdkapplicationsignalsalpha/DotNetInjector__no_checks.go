//go:build no_runtime_type_checking

package awscdkapplicationsignalsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DotNetInjector) validateInjectAdditionalEnvironmentsParameters(envsToInject *map[string]*string, envsFromTaskDef *map[string]*string) error {
	return nil
}

func (d *jsiiProxy_DotNetInjector) validateInjectInitContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (d *jsiiProxy_DotNetInjector) validateOverrideAdditionalEnvironmentsParameters(envsToOverride *map[string]*string, envsFromTaskDef *map[string]*string) error {
	return nil
}

func (d *jsiiProxy_DotNetInjector) validateRenderDefaultContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (j *jsiiProxy_DotNetInjector) validateSetInstrumentationVersionParameters(val InstrumentationVersion) error {
	return nil
}

func (j *jsiiProxy_DotNetInjector) validateSetSharedVolumeNameParameters(val *string) error {
	return nil
}

func validateNewDotNetInjectorParameters(sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) error {
	return nil
}


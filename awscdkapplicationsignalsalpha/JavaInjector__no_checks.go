//go:build no_runtime_type_checking

package awscdkapplicationsignalsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JavaInjector) validateInjectAdditionalEnvironmentsParameters(envsToInject *map[string]*string, _envsFromTaskDef *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_JavaInjector) validateInjectInitContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (j *jsiiProxy_JavaInjector) validateOverrideAdditionalEnvironmentsParameters(_envsToOverride *map[string]*string, _overrideEnvironments *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_JavaInjector) validateRenderDefaultContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (j *jsiiProxy_JavaInjector) validateSetInstrumentationVersionParameters(val InstrumentationVersion) error {
	return nil
}

func (j *jsiiProxy_JavaInjector) validateSetSharedVolumeNameParameters(val *string) error {
	return nil
}

func validateNewJavaInjectorParameters(sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) error {
	return nil
}


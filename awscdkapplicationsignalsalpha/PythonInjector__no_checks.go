//go:build no_runtime_type_checking

package awscdkapplicationsignalsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PythonInjector) validateInjectAdditionalEnvironmentsParameters(envsToInject *map[string]*string, _envsFromTaskDef *map[string]*string) error {
	return nil
}

func (p *jsiiProxy_PythonInjector) validateInjectInitContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (p *jsiiProxy_PythonInjector) validateOverrideAdditionalEnvironmentsParameters(envsToOverride *map[string]*string, envsFromTaskDef *map[string]*string) error {
	return nil
}

func (p *jsiiProxy_PythonInjector) validateRenderDefaultContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (j *jsiiProxy_PythonInjector) validateSetInstrumentationVersionParameters(val InstrumentationVersion) error {
	return nil
}

func (j *jsiiProxy_PythonInjector) validateSetSharedVolumeNameParameters(val *string) error {
	return nil
}

func validateNewPythonInjectorParameters(sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) error {
	return nil
}


//go:build no_runtime_type_checking

package awscdkapplicationsignalsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_Injector) validateInjectAdditionalEnvironmentsParameters(envsToInject *map[string]*string, envsFromTaskDef *map[string]*string) error {
	return nil
}

func (i *jsiiProxy_Injector) validateInjectInitContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (i *jsiiProxy_Injector) validateOverrideAdditionalEnvironmentsParameters(envsToOverride *map[string]*string, envsFromTaskDef *map[string]*string) error {
	return nil
}

func (i *jsiiProxy_Injector) validateRenderDefaultContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (j *jsiiProxy_Injector) validateSetInstrumentationVersionParameters(val InstrumentationVersion) error {
	return nil
}

func (j *jsiiProxy_Injector) validateSetSharedVolumeNameParameters(val *string) error {
	return nil
}

func validateNewInjectorParameters(sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) error {
	return nil
}


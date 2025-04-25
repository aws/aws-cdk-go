//go:build !no_runtime_type_checking

package awscdkapplicationsignalsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

func (j *jsiiProxy_JavaInjector) validateInjectAdditionalEnvironmentsParameters(envsToInject *map[string]*string, _envsFromTaskDef *map[string]*string) error {
	if envsToInject == nil {
		return fmt.Errorf("parameter envsToInject is required, but nil was provided")
	}

	if _envsFromTaskDef == nil {
		return fmt.Errorf("parameter _envsFromTaskDef is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JavaInjector) validateInjectInitContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	if taskDefinition == nil {
		return fmt.Errorf("parameter taskDefinition is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JavaInjector) validateOverrideAdditionalEnvironmentsParameters(_envsToOverride *map[string]*string, _overrideEnvironments *map[string]*string) error {
	if _envsToOverride == nil {
		return fmt.Errorf("parameter _envsToOverride is required, but nil was provided")
	}

	if _overrideEnvironments == nil {
		return fmt.Errorf("parameter _overrideEnvironments is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JavaInjector) validateRenderDefaultContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	if taskDefinition == nil {
		return fmt.Errorf("parameter taskDefinition is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JavaInjector) validateSetInstrumentationVersionParameters(val InstrumentationVersion) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JavaInjector) validateSetSharedVolumeNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewJavaInjectorParameters(sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) error {
	if sharedVolumeName == nil {
		return fmt.Errorf("parameter sharedVolumeName is required, but nil was provided")
	}

	if instrumentationVersion == nil {
		return fmt.Errorf("parameter instrumentationVersion is required, but nil was provided")
	}

	for idx_f63c31, v := range *overrideEnvironments {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter overrideEnvironments[%#v]", idx_f63c31) }); err != nil {
			return err
		}
	}

	return nil
}


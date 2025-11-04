//go:build !no_runtime_type_checking

package awscdkapplicationsignalsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

func (n *jsiiProxy_NodeInjector) validateInjectAdditionalEnvironmentsParameters(envsToInject *map[string]*string, envsFromTaskDef *map[string]*string) error {
	if envsToInject == nil {
		return fmt.Errorf("parameter envsToInject is required, but nil was provided")
	}

	if envsFromTaskDef == nil {
		return fmt.Errorf("parameter envsFromTaskDef is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NodeInjector) validateInjectInitContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	if taskDefinition == nil {
		return fmt.Errorf("parameter taskDefinition is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NodeInjector) validateOverrideAdditionalEnvironmentsParameters(envsToOverride *map[string]*string, envsFromTaskDef *map[string]*string) error {
	if envsToOverride == nil {
		return fmt.Errorf("parameter envsToOverride is required, but nil was provided")
	}

	if envsFromTaskDef == nil {
		return fmt.Errorf("parameter envsFromTaskDef is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NodeInjector) validateRenderDefaultContainerParameters(taskDefinition awsecs.TaskDefinition) error {
	if taskDefinition == nil {
		return fmt.Errorf("parameter taskDefinition is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NodeInjector) validateSetInstrumentationVersionParameters(val InstrumentationVersion) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NodeInjector) validateSetSharedVolumeNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewNodeInjectorParameters(sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) error {
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


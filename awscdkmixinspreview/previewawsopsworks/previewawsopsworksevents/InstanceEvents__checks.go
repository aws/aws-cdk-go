//go:build !no_runtime_type_checking

package previewawsopsworksevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsopsworks"
)

func (i *jsiiProxy_InstanceEvents) validateOpsWorksAlertPatternParameters(options *InstanceEvents_OpsWorksAlert_OpsWorksAlertProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_InstanceEvents) validateOpsWorksCommandStateChangePatternParameters(options *InstanceEvents_OpsWorksCommandStateChange_OpsWorksCommandStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_InstanceEvents) validateOpsWorksInstanceStateChangePatternParameters(options *InstanceEvents_OpsWorksInstanceStateChange_OpsWorksInstanceStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInstanceEvents_FromInstanceParameters(instanceRef interfacesawsopsworks.IInstanceRef) error {
	if instanceRef == nil {
		return fmt.Errorf("parameter instanceRef is required, but nil was provided")
	}

	return nil
}


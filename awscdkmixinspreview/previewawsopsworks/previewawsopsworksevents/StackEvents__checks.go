//go:build !no_runtime_type_checking

package previewawsopsworksevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsopsworks"
)

func (s *jsiiProxy_StackEvents) validateOpsWorksDeploymentStateChangePatternParameters(options *StackEvents_OpsWorksDeploymentStateChange_OpsWorksDeploymentStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateStackEvents_FromStackParameters(stackRef interfacesawsopsworks.IStackRef) error {
	if stackRef == nil {
		return fmt.Errorf("parameter stackRef is required, but nil was provided")
	}

	return nil
}


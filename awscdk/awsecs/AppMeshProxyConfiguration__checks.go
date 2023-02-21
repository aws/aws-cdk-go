//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AppMeshProxyConfiguration) validateBindParameters(_scope constructs.Construct, _taskDefinition TaskDefinition) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _taskDefinition == nil {
		return fmt.Errorf("parameter _taskDefinition is required, but nil was provided")
	}

	return nil
}

func validateNewAppMeshProxyConfigurationParameters(props *AppMeshProxyConfigurationConfigProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


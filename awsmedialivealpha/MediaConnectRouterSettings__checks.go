//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateMediaConnectRouterSettings_PerPipelineParameters(settings *MediaConnectRouterPerPipelineSettings) error {
	if settings == nil {
		return fmt.Errorf("parameter settings is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(settings, func() string { return "parameter settings" }); err != nil {
		return err
	}

	return nil
}

func validateMediaConnectRouterSettings_SharedParameters(settings *MediaConnectRouterPipelineConfig) error {
	if err := _jsii_.ValidateStruct(settings, func() string { return "parameter settings" }); err != nil {
		return err
	}

	return nil
}


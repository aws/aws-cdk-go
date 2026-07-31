//go:build !no_runtime_type_checking

package awsmediaconnectalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateOutputConfiguration_NdiParameters(input *NdiOutputConfig) error {
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}

func validateOutputConfiguration_RistParameters(input *RistOutputConfig) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}

func validateOutputConfiguration_RouterParameters(input *RouterTransitConfig) error {
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}

func validateOutputConfiguration_RtpParameters(input *RtpOutputConfig) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}

func validateOutputConfiguration_RtpFecParameters(input *RtpFecOutputConfig) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}

func validateOutputConfiguration_SrtCallerParameters(input *SrtCallerOutputConfig) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}

func validateOutputConfiguration_SrtListenerParameters(input *SrtListenerOutputConfig) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}

func validateOutputConfiguration_ZixiPullParameters(input *ZixiPullOutputConfig) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}

func validateOutputConfiguration_ZixiPushParameters(input *ZixiPushOutputConfig) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}


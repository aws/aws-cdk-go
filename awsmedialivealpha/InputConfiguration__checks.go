//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateInputConfiguration_CdiParameters(props *CdiInputProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateInputConfiguration_InputDeviceParameters(props *InputDeviceInputProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateInputConfiguration_MediaConnectParameters(props *MediaConnectInputProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateInputConfiguration_MediaConnectRouterParameters(props *MediaConnectRouterInputProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateInputConfiguration_Mp4FileParameters(sources *[]InputSource) error {
	if sources == nil {
		return fmt.Errorf("parameter sources is required, but nil was provided")
	}

	return nil
}

func validateInputConfiguration_MulticastParameters(props *MulticastInputProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateInputConfiguration_RtmpPullParameters(sources *[]InputSource) error {
	if sources == nil {
		return fmt.Errorf("parameter sources is required, but nil was provided")
	}

	return nil
}

func validateInputConfiguration_RtmpPushParameters(props *PushInputProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateInputConfiguration_RtpPushParameters(props *PushInputProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateInputConfiguration_SdiParameters(sources *[]ISdiSource) error {
	if sources == nil {
		return fmt.Errorf("parameter sources is required, but nil was provided")
	}

	return nil
}

func validateInputConfiguration_Smpte2110ReceiverGroupParameters(props *Smpte2110InputProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateInputConfiguration_SrtCallerParameters(sources *[]*SrtCallerSourceProps) error {
	if sources == nil {
		return fmt.Errorf("parameter sources is required, but nil was provided")
	}
	for idx_878a52, v := range *sources {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter sources[%#v]", idx_878a52) }); err != nil {
			return err
		}
	}

	return nil
}

func validateInputConfiguration_SrtListenerParameters(props *SrtListenerInputProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateInputConfiguration_TsFileParameters(sources *[]InputSource) error {
	if sources == nil {
		return fmt.Errorf("parameter sources is required, but nil was provided")
	}

	return nil
}

func validateInputConfiguration_UdpPushParameters(props *PushInputProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateInputConfiguration_UrlPullParameters(sources *[]InputSource) error {
	if sources == nil {
		return fmt.Errorf("parameter sources is required, but nil was provided")
	}

	return nil
}


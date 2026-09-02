//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateInputConfiguration_CdiParameters(props *CdiInputProps) error {
	return nil
}

func validateInputConfiguration_InputDeviceParameters(props *InputDeviceInputProps) error {
	return nil
}

func validateInputConfiguration_MediaConnectParameters(props *MediaConnectInputProps) error {
	return nil
}

func validateInputConfiguration_MediaConnectRouterParameters(props *MediaConnectRouterInputProps) error {
	return nil
}

func validateInputConfiguration_Mp4FileParameters(sources *[]InputSource) error {
	return nil
}

func validateInputConfiguration_MulticastParameters(props *MulticastInputProps) error {
	return nil
}

func validateInputConfiguration_RtmpPullParameters(sources *[]InputSource) error {
	return nil
}

func validateInputConfiguration_RtmpPushParameters(props *PushInputProps) error {
	return nil
}

func validateInputConfiguration_RtpPushParameters(props *PushInputProps) error {
	return nil
}

func validateInputConfiguration_SdiParameters(sources *[]ISdiSource) error {
	return nil
}

func validateInputConfiguration_Smpte2110ReceiverGroupParameters(props *Smpte2110InputProps) error {
	return nil
}

func validateInputConfiguration_SrtCallerParameters(sources *[]*SrtCallerSourceProps) error {
	return nil
}

func validateInputConfiguration_SrtListenerParameters(props *SrtListenerInputProps) error {
	return nil
}

func validateInputConfiguration_TsFileParameters(sources *[]InputSource) error {
	return nil
}

func validateInputConfiguration_UdpPushParameters(props *PushInputProps) error {
	return nil
}

func validateInputConfiguration_UrlPullParameters(sources *[]InputSource) error {
	return nil
}


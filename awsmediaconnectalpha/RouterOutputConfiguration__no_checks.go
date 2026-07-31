//go:build no_runtime_type_checking

package awsmediaconnectalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateRouterOutputConfiguration_MediaConnectFlowParameters(props *MediaConnectFlowConnectionProps) error {
	return nil
}

func validateRouterOutputConfiguration_MediaConnectFlowWithoutConnectionParameters(props *MediaConnectFlowNoConnectionProps) error {
	return nil
}

func validateRouterOutputConfiguration_MediaLiveInputParameters(props *MediaLiveInputConnectionProps) error {
	return nil
}

func validateRouterOutputConfiguration_MediaLiveInputWithoutConnectionParameters(props *MediaLiveNoInputConnectionProps) error {
	return nil
}

func validateRouterOutputConfiguration_StandardParameters(props *StandardOutputConfigurationProps) error {
	return nil
}


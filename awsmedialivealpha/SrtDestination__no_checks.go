//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateSrtDestination_CallerParameters(props *SrtCallerDestinationProps) error {
	return nil
}

func validateSrtDestination_CallerUrlParameters(url *string, options *SrtCallerUrlOptions) error {
	return nil
}

func validateSrtDestination_ListenerParameters(props *SrtListenerDestinationProps) error {
	return nil
}


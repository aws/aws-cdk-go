//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateUdpOutputDestination_RtpParameters(props *TransportOutputDestinationProps) error {
	return nil
}

func validateUdpOutputDestination_UdpParameters(props *TransportOutputDestinationProps) error {
	return nil
}

func validateUdpOutputDestination_UrlParameters(url *string) error {
	return nil
}


//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateRtmpDestination_UrlParameters(url *string, streamName *string, options *OutputDestinationOptions) error {
	return nil
}


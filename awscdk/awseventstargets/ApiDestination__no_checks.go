//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiDestination) validateBindParameters(rule interfacesawsevents.IRuleRef) error {
	return nil
}

func validateNewApiDestinationParameters(apiDestination awsevents.IApiDestination, props *ApiDestinationProps) error {
	return nil
}


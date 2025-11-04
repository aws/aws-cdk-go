//go:build no_runtime_type_checking

package awscdkpipesenrichmentsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiDestinationEnrichment) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (a *jsiiProxy_ApiDestinationEnrichment) validateGrantInvokeParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewApiDestinationEnrichmentParameters(destination awsevents.IApiDestination, props *ApiDestinationEnrichmentProps) error {
	return nil
}


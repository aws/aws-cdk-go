//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualGatewayGrants) validateStreamAggregatedResourcesParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateVirtualGatewayGrants_FromVirtualGatewayParameters(resource interfacesawsappmesh.IVirtualGatewayRef) error {
	return nil
}


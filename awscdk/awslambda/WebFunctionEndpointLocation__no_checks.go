//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func validateWebFunctionEndpointLocation_MultiRegionParameters(replicaRegions *[]*string) error {
	return nil
}

func validateWebFunctionEndpointLocation_PerRegionParameters(replicaRegions *[]*string) error {
	return nil
}


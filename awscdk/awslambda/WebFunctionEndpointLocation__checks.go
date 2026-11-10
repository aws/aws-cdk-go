//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func validateWebFunctionEndpointLocation_MultiRegionParameters(replicaRegions *[]*string) error {
	if replicaRegions == nil {
		return fmt.Errorf("parameter replicaRegions is required, but nil was provided")
	}

	return nil
}

func validateWebFunctionEndpointLocation_PerRegionParameters(replicaRegions *[]*string) error {
	if replicaRegions == nil {
		return fmt.Errorf("parameter replicaRegions is required, but nil was provided")
	}

	return nil
}


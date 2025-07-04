//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"
)

func validateClientVpnRouteTarget_SubnetParameters(subnet ISubnet) error {
	if subnet == nil {
		return fmt.Errorf("parameter subnet is required, but nil was provided")
	}

	return nil
}


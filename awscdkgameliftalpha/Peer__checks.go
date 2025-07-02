//go:build !no_runtime_type_checking

package awscdkgameliftalpha

import (
	"fmt"
)

func validatePeer_Ipv4Parameters(cidrIp *string) error {
	if cidrIp == nil {
		return fmt.Errorf("parameter cidrIp is required, but nil was provided")
	}

	return nil
}


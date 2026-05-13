//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"
)

func (i *jsiiProxy_ICredentialProviderConfig) validateGrantNeededPermissionsToRoleParameters(gateway IGateway) error {
	if gateway == nil {
		return fmt.Errorf("parameter gateway is required, but nil was provided")
	}

	return nil
}


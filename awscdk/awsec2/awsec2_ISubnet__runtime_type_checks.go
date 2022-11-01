//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"
)

func (i *jsiiProxy_ISubnet) validateAssociateNetworkAclParameters(id *string, acl INetworkAcl) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if acl == nil {
		return fmt.Errorf("parameter acl is required, but nil was provided")
	}

	return nil
}


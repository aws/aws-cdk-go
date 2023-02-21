//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (n *jsiiProxy_NatInstanceImage) validateGetImageParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}


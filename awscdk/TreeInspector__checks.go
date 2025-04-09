//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func (t *jsiiProxy_TreeInspector) validateAddAttributeParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}


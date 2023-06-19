//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func (c *jsiiProxy_ConcreteDependable) validateAddParameters(construct IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}


//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func (j *jsiiProxy_JsonNull) validateResolveParameters(_ctx IResolveContext) error {
	if _ctx == nil {
		return fmt.Errorf("parameter _ctx is required, but nil was provided")
	}

	return nil
}


//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateShims_AsAspectParameters(mixin constructs.IMixin) error {
	if mixin == nil {
		return fmt.Errorf("parameter mixin is required, but nil was provided")
	}

	return nil
}

func validateShims_AsMixinParameters(aspect IAspect) error {
	if aspect == nil {
		return fmt.Errorf("parameter aspect is required, but nil was provided")
	}

	return nil
}


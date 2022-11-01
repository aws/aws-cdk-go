//go:build !no_runtime_type_checking

package regioninfo

import (
	"fmt"
)

func validateFact_FindParameters(region *string, name *string) error {
	if region == nil {
		return fmt.Errorf("parameter region is required, but nil was provided")
	}

	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateFact_RegisterParameters(fact IFact) error {
	if fact == nil {
		return fmt.Errorf("parameter fact is required, but nil was provided")
	}

	return nil
}

func validateFact_RequireFactParameters(region *string, name *string) error {
	if region == nil {
		return fmt.Errorf("parameter region is required, but nil was provided")
	}

	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateFact_UnregisterParameters(region *string, name *string) error {
	if region == nil {
		return fmt.Errorf("parameter region is required, but nil was provided")
	}

	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}


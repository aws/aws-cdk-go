//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"
)

func validateVideoSelection_ByPidParameters(pid *float64) error {
	if pid == nil {
		return fmt.Errorf("parameter pid is required, but nil was provided")
	}

	return nil
}

func validateVideoSelection_ByProgramIdParameters(programId *float64) error {
	if programId == nil {
		return fmt.Errorf("parameter programId is required, but nil was provided")
	}

	return nil
}


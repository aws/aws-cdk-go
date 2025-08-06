//go:build !no_runtime_type_checking

package awscdkivsalpha

import (
	"fmt"
)

func validateRenditionConfiguration_CustomParameters(renditions *[]Resolution) error {
	if renditions == nil {
		return fmt.Errorf("parameter renditions is required, but nil was provided")
	}

	return nil
}


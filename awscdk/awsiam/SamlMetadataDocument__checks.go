//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func validateSamlMetadataDocument_FromFileParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateSamlMetadataDocument_FromXmlParameters(xml *string) error {
	if xml == nil {
		return fmt.Errorf("parameter xml is required, but nil was provided")
	}

	return nil
}


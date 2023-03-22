//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"
)

func (j *jsiiProxy_ISchemaConfig) validateSetApiIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ISchemaConfig) validateSetDefinitionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}


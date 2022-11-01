//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"
)

func (i *jsiiProxy_ITaskDefinitionExtension) validateExtendParameters(taskDefinition TaskDefinition) error {
	if taskDefinition == nil {
		return fmt.Errorf("parameter taskDefinition is required, but nil was provided")
	}

	return nil
}


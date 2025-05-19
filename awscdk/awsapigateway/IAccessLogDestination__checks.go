//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"
)

func (i *jsiiProxy_IAccessLogDestination) validateBindParameters(stage IStage) error {
	if stage == nil {
		return fmt.Errorf("parameter stage is required, but nil was provided")
	}

	return nil
}


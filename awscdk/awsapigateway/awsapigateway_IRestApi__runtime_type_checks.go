//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"
)

func (j *jsiiProxy_IRestApi) validateSetDeploymentStageParameters(val Stage) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}


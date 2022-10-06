//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awscodedeploy

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v3"
)

func validateEcsDeploymentConfig_FromEcsDeploymentConfigNameParameters(_scope constructs.Construct, _id *string, ecsDeploymentConfigName *string) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _id == nil {
		return fmt.Errorf("parameter _id is required, but nil was provided")
	}

	if ecsDeploymentConfigName == nil {
		return fmt.Errorf("parameter ecsDeploymentConfigName is required, but nil was provided")
	}

	return nil
}


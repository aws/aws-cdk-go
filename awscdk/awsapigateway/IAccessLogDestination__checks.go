//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway"
)

func (i *jsiiProxy_IAccessLogDestination) validateBindParameters(stage interfacesawsapigateway.IStageRef) error {
	if stage == nil {
		return fmt.Errorf("parameter stage is required, but nil was provided")
	}

	return nil
}


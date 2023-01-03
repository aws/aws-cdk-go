//go:build !no_runtime_type_checking

package awsroute53targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
)

func (e *jsiiProxy_ElasticBeanstalkEnvironmentEndpointTarget) validateBindParameters(_record awsroute53.IRecordSet) error {
	if _record == nil {
		return fmt.Errorf("parameter _record is required, but nil was provided")
	}

	return nil
}

func validateNewElasticBeanstalkEnvironmentEndpointTargetParameters(environmentEndpoint *string) error {
	if environmentEndpoint == nil {
		return fmt.Errorf("parameter environmentEndpoint is required, but nil was provided")
	}

	return nil
}


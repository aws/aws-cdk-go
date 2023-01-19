//go:build !no_runtime_type_checking

package awsroute53targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
)

func (a *jsiiProxy_ApiGateway) validateBindParameters(_record awsroute53.IRecordSet) error {
	if _record == nil {
		return fmt.Errorf("parameter _record is required, but nil was provided")
	}

	return nil
}

func validateNewApiGatewayParameters(api awsapigateway.RestApiBase) error {
	if api == nil {
		return fmt.Errorf("parameter api is required, but nil was provided")
	}

	return nil
}


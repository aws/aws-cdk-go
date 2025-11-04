//go:build !no_runtime_type_checking

package awsroute53targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

func (a *jsiiProxy_ApiGatewayDomain) validateBindParameters(record awsroute53.IRecordSet) error {
	if record == nil {
		return fmt.Errorf("parameter record is required, but nil was provided")
	}

	return nil
}

func validateNewApiGatewayDomainParameters(domainName awsapigateway.IDomainName) error {
	if domainName == nil {
		return fmt.Errorf("parameter domainName is required, but nil was provided")
	}

	return nil
}


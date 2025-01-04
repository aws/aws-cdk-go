//go:build !no_runtime_type_checking

package awsroute53targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

func (u *jsiiProxy_UserPoolDomainTarget) validateBindParameters(record awsroute53.IRecordSet) error {
	if record == nil {
		return fmt.Errorf("parameter record is required, but nil was provided")
	}

	return nil
}

func validateNewUserPoolDomainTargetParameters(domain awscognito.UserPoolDomain) error {
	if domain == nil {
		return fmt.Errorf("parameter domain is required, but nil was provided")
	}

	return nil
}


//go:build !no_runtime_type_checking

package awsroute53targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

func (a *jsiiProxy_ApiGatewayv2DomainProperties) validateBindParameters(record awsroute53.IRecordSet) error {
	if record == nil {
		return fmt.Errorf("parameter record is required, but nil was provided")
	}

	return nil
}

func validateNewApiGatewayv2DomainPropertiesParameters(regionalDomainName *string, regionalHostedZoneId *string) error {
	if regionalDomainName == nil {
		return fmt.Errorf("parameter regionalDomainName is required, but nil was provided")
	}

	if regionalHostedZoneId == nil {
		return fmt.Errorf("parameter regionalHostedZoneId is required, but nil was provided")
	}

	return nil
}


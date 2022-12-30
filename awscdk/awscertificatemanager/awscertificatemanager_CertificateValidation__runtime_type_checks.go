//go:build !no_runtime_type_checking

package awscertificatemanager

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
)

func validateCertificateValidation_FromDnsMultiZoneParameters(hostedZones *map[string]awsroute53.IHostedZone) error {
	if hostedZones == nil {
		return fmt.Errorf("parameter hostedZones is required, but nil was provided")
	}

	return nil
}


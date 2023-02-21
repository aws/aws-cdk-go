//go:build !no_runtime_type_checking

package awsses

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

func validateIdentity_DomainParameters(domain *string) error {
	if domain == nil {
		return fmt.Errorf("parameter domain is required, but nil was provided")
	}

	return nil
}

func validateIdentity_EmailParameters(email *string) error {
	if email == nil {
		return fmt.Errorf("parameter email is required, but nil was provided")
	}

	return nil
}

func validateIdentity_PublicHostedZoneParameters(hostedZone awsroute53.IPublicHostedZone) error {
	if hostedZone == nil {
		return fmt.Errorf("parameter hostedZone is required, but nil was provided")
	}

	return nil
}


//go:build !no_runtime_type_checking

package awselasticloadbalancing

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

func validateNewListenerPortParameters(securityGroup awsec2.ISecurityGroup, defaultPort awsec2.Port) error {
	if securityGroup == nil {
		return fmt.Errorf("parameter securityGroup is required, but nil was provided")
	}

	if defaultPort == nil {
		return fmt.Errorf("parameter defaultPort is required, but nil was provided")
	}

	return nil
}


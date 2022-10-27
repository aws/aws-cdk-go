//go:build !no_runtime_type_checking

package awsapprunner

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (i *jsiiProxy_IVpcConnector) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}


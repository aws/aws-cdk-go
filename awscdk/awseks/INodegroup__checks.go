//go:build !no_runtime_type_checking

package awseks

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (i *jsiiProxy_INodegroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}


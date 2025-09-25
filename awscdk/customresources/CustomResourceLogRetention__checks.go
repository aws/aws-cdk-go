//go:build !no_runtime_type_checking

package customresources

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CustomResourceLogRetention) validateVisitParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func validateNewCustomResourceLogRetentionParameters(setLogRetention awslogs.RetentionDays) error {
	if setLogRetention == "" {
		return fmt.Errorf("parameter setLogRetention is required, but nil was provided")
	}

	return nil
}


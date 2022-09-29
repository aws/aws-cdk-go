//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (f *jsiiProxy_FlowLogDestination) validateBindParameters(scope awscdk.Construct, flowLog FlowLog) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if flowLog == nil {
		return fmt.Errorf("parameter flowLog is required, but nil was provided")
	}

	return nil
}


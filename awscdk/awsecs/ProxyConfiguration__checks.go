//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (p *jsiiProxy_ProxyConfiguration) validateBindParameters(_scope constructs.Construct, _taskDefinition TaskDefinition) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _taskDefinition == nil {
		return fmt.Errorf("parameter _taskDefinition is required, but nil was provided")
	}

	return nil
}


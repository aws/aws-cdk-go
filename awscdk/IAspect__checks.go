//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_IAspect) validateVisitParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}


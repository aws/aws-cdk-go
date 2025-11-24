//go:build !no_runtime_type_checking

package mixinsawss3

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_EnableVersioning) validateApplyToParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EnableVersioning) validateSupportsParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}


//go:build !no_runtime_type_checking

package previewawss3mixins

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AutoDeleteObjects) validateApplyToParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AutoDeleteObjects) validateSupportsParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}


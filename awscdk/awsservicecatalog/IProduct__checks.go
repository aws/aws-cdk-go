//go:build !no_runtime_type_checking

package awsservicecatalog

import (
	"fmt"
)

func (i *jsiiProxy_IProduct) validateAssociateTagOptionsParameters(tagOptions TagOptions) error {
	if tagOptions == nil {
		return fmt.Errorf("parameter tagOptions is required, but nil was provided")
	}

	return nil
}


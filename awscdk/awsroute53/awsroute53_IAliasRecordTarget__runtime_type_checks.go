//go:build !no_runtime_type_checking

package awsroute53

import (
	"fmt"
)

func (i *jsiiProxy_IAliasRecordTarget) validateBindParameters(record IRecordSet) error {
	if record == nil {
		return fmt.Errorf("parameter record is required, but nil was provided")
	}

	return nil
}


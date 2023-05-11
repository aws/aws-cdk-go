//go:build !no_runtime_type_checking

package awsses

import (
	"fmt"
)

func (i *jsiiProxy_IReceiptRuleAction) validateBindParameters(receiptRule IReceiptRule) error {
	if receiptRule == nil {
		return fmt.Errorf("parameter receiptRule is required, but nil was provided")
	}

	return nil
}


//go:build !no_runtime_type_checking

package awsses

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses"
)

func (i *jsiiProxy_IReceiptRuleAction) validateBindParameters(receiptRule interfacesawsses.IReceiptRuleRef) error {
	if receiptRule == nil {
		return fmt.Errorf("parameter receiptRule is required, but nil was provided")
	}

	return nil
}


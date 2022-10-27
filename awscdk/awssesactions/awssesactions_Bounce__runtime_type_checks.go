//go:build !no_runtime_type_checking

package awssesactions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsses"
)

func (b *jsiiProxy_Bounce) validateBindParameters(_rule awsses.IReceiptRule) error {
	if _rule == nil {
		return fmt.Errorf("parameter _rule is required, but nil was provided")
	}

	return nil
}

func validateNewBounceParameters(props *BounceProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


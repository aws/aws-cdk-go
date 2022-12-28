//go:build !no_runtime_type_checking

package awssnssubscriptions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

func (s *jsiiProxy_SmsSubscription) validateBindParameters(_topic awssns.ITopic) error {
	if _topic == nil {
		return fmt.Errorf("parameter _topic is required, but nil was provided")
	}

	return nil
}

func validateNewSmsSubscriptionParameters(phoneNumber *string, props *SmsSubscriptionProps) error {
	if phoneNumber == nil {
		return fmt.Errorf("parameter phoneNumber is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


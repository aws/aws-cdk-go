//go:build !no_runtime_type_checking

package awssnssubscriptions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

func (u *jsiiProxy_UrlSubscription) validateBindParameters(_topic awssns.ITopic) error {
	if _topic == nil {
		return fmt.Errorf("parameter _topic is required, but nil was provided")
	}

	return nil
}

func validateNewUrlSubscriptionParameters(url *string, props *UrlSubscriptionProps) error {
	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


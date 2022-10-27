//go:build !no_runtime_type_checking

package awseventstargets

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

func (a *jsiiProxy_AwsApi) validateBindParameters(rule awsevents.IRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}

	return nil
}

func validateNewAwsApiParameters(props *AwsApiProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


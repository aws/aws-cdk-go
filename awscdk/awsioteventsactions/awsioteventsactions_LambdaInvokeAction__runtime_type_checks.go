//go:build !no_runtime_type_checking

package awsioteventsactions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiotevents"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
)

func (l *jsiiProxy_LambdaInvokeAction) validateBindParameters(_scope constructs.Construct, options *awsiotevents.ActionBindOptions) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewLambdaInvokeActionParameters(func_ awslambda.IFunction) error {
	if func_ == nil {
		return fmt.Errorf("parameter func_ is required, but nil was provided")
	}

	return nil
}


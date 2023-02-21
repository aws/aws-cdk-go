//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (f *jsiiProxy_FlowLogDestination) validateBindParameters(scope constructs.Construct, flowLog FlowLog) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if flowLog == nil {
		return fmt.Errorf("parameter flowLog is required, but nil was provided")
	}

	return nil
}

func validateFlowLogDestination_ToS3Parameters(options *S3DestinationOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


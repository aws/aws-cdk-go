//go:build !no_runtime_type_checking

package awslambdaeventsources

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

func (m *jsiiProxy_ManagedKafkaEventSource) validateBindParameters(target awslambda.IFunction) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedKafkaEventSource) validateEnrichMappingOptionsParameters(options *awslambda.EventSourceMappingOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewManagedKafkaEventSourceParameters(props *ManagedKafkaEventSourceProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


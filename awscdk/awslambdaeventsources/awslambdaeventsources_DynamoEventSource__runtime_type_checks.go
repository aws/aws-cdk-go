//go:build !no_runtime_type_checking

package awslambdaeventsources

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

func (d *jsiiProxy_DynamoEventSource) validateBindParameters(target awslambda.IFunction) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DynamoEventSource) validateEnrichMappingOptionsParameters(options *awslambda.EventSourceMappingOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewDynamoEventSourceParameters(table awsdynamodb.ITable, props *DynamoEventSourceProps) error {
	if table == nil {
		return fmt.Errorf("parameter table is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


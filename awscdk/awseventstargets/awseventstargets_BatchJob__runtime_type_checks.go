//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awseventstargets

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

func (b *jsiiProxy_BatchJob) validateBindParameters(rule awsevents.IRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}

	return nil
}

func validateNewBatchJobParameters(jobQueueArn *string, jobQueueScope awscdk.IConstruct, jobDefinitionArn *string, jobDefinitionScope awscdk.IConstruct, props *BatchJobProps) error {
	if jobQueueArn == nil {
		return fmt.Errorf("parameter jobQueueArn is required, but nil was provided")
	}

	if jobQueueScope == nil {
		return fmt.Errorf("parameter jobQueueScope is required, but nil was provided")
	}

	if jobDefinitionArn == nil {
		return fmt.Errorf("parameter jobDefinitionArn is required, but nil was provided")
	}

	if jobDefinitionScope == nil {
		return fmt.Errorf("parameter jobDefinitionScope is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


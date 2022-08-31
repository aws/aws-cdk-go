//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

func (r *jsiiProxy_RunGlueJobTask) validateBindParameters(task awsstepfunctions.Task) error {
	if task == nil {
		return fmt.Errorf("parameter task is required, but nil was provided")
	}

	return nil
}

func validateNewRunGlueJobTaskParameters(glueJobName *string, props *RunGlueJobTaskProps) error {
	if glueJobName == nil {
		return fmt.Errorf("parameter glueJobName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


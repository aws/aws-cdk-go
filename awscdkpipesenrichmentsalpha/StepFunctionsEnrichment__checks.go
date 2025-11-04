//go:build !no_runtime_type_checking

package awscdkpipesenrichmentsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

func (s *jsiiProxy_StepFunctionsEnrichment) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	if pipe == nil {
		return fmt.Errorf("parameter pipe is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StepFunctionsEnrichment) validateGrantInvokeParameters(grantee awsiam.IRole) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateNewStepFunctionsEnrichmentParameters(stateMachine awsstepfunctions.IStateMachine, props *StepFunctionsEnrichmentProps) error {
	if stateMachine == nil {
		return fmt.Errorf("parameter stateMachine is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


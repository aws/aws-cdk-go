//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/constructs-go/constructs/v3"
)

func (u *jsiiProxy_UpdatePipelineAction) validateBindParameters(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if stage == nil {
		return fmt.Errorf("parameter stage is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (u *jsiiProxy_UpdatePipelineAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (u *jsiiProxy_UpdatePipelineAction) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (u *jsiiProxy_UpdatePipelineAction) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func validateUpdatePipelineAction_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewUpdatePipelineActionParameters(scope constructs.Construct, id *string, props *UpdatePipelineActionProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


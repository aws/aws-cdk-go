//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awscodepipelineactions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

func (j *jsiiProxy_JenkinsAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
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

func (j *jsiiProxy_JenkinsAction) validateBoundParameters(_scope constructs.Construct, _stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _stage == nil {
		return fmt.Errorf("parameter _stage is required, but nil was provided")
	}

	if _options == nil {
		return fmt.Errorf("parameter _options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_JenkinsAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_JenkinsAction) validateVariableExpressionParameters(variableName *string) error {
	if variableName == nil {
		return fmt.Errorf("parameter variableName is required, but nil was provided")
	}

	return nil
}

func validateNewJenkinsActionParameters(props *JenkinsActionProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


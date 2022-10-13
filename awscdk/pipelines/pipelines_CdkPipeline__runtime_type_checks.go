//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package pipelines

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
)

func (c *jsiiProxy_CdkPipeline) validateAddApplicationStageParameters(appStage awscdk.Stage, options *AddStageOptions) error {
	if appStage == nil {
		return fmt.Errorf("parameter appStage is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CdkPipeline) validateAddStageParameters(stageName *string, options *BaseStageOptions) error {
	if stageName == nil {
		return fmt.Errorf("parameter stageName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CdkPipeline) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdkPipeline) validateStackOutputParameters(cfnOutput awscdk.CfnOutput) error {
	if cfnOutput == nil {
		return fmt.Errorf("parameter cfnOutput is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdkPipeline) validateStageParameters(stageName *string) error {
	if stageName == nil {
		return fmt.Errorf("parameter stageName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdkPipeline) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func validateCdkPipeline_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewCdkPipelineParameters(scope constructs.Construct, id *string, props *CdkPipelineProps) error {
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


//go:build !no_runtime_type_checking

package awscdkschedulertargetsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
)

func (c *jsiiProxy_CodePipelineStartPipelineExecution) validateAddTargetActionToRoleParameters(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole) error {
	if schedule == nil {
		return fmt.Errorf("parameter schedule is required, but nil was provided")
	}

	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CodePipelineStartPipelineExecution) validateBindParameters(schedule awscdkscheduleralpha.ISchedule) error {
	if schedule == nil {
		return fmt.Errorf("parameter schedule is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CodePipelineStartPipelineExecution) validateBindBaseTargetConfigParameters(_schedule awscdkscheduleralpha.ISchedule) error {
	if _schedule == nil {
		return fmt.Errorf("parameter _schedule is required, but nil was provided")
	}

	return nil
}

func validateNewCodePipelineStartPipelineExecutionParameters(pipeline awscodepipeline.IPipeline, props *ScheduleTargetBaseProps) error {
	if pipeline == nil {
		return fmt.Errorf("parameter pipeline is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


//go:build !no_runtime_type_checking

package awscdkschedulertargetsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
)

func (s *jsiiProxy_ScheduleTargetBase) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ScheduleTargetBase) validateBindParameters(schedule awscdkscheduleralpha.ISchedule) error {
	if schedule == nil {
		return fmt.Errorf("parameter schedule is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ScheduleTargetBase) validateBindBaseTargetConfigParameters(_schedule awscdkscheduleralpha.ISchedule) error {
	if _schedule == nil {
		return fmt.Errorf("parameter _schedule is required, but nil was provided")
	}

	return nil
}

func validateNewScheduleTargetBaseParameters(baseProps *ScheduleTargetBaseProps, targetArn *string) error {
	if baseProps == nil {
		return fmt.Errorf("parameter baseProps is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(baseProps, func() string { return "parameter baseProps" }); err != nil {
		return err
	}

	if targetArn == nil {
		return fmt.Errorf("parameter targetArn is required, but nil was provided")
	}

	return nil
}


//go:build !no_runtime_type_checking

package awscdkschedulertargetsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
)

func (u *jsiiProxy_Universal) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

func (u *jsiiProxy_Universal) validateBindParameters(schedule awscdkscheduleralpha.ISchedule) error {
	if schedule == nil {
		return fmt.Errorf("parameter schedule is required, but nil was provided")
	}

	return nil
}

func (u *jsiiProxy_Universal) validateBindBaseTargetConfigParameters(_schedule awscdkscheduleralpha.ISchedule) error {
	if _schedule == nil {
		return fmt.Errorf("parameter _schedule is required, but nil was provided")
	}

	return nil
}

func validateNewUniversalParameters(props *UniversalTargetProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


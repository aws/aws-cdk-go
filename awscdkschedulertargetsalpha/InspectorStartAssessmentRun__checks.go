//go:build !no_runtime_type_checking

package awscdkschedulertargetsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspector"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
)

func (i *jsiiProxy_InspectorStartAssessmentRun) validateAddTargetActionToRoleParameters(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole) error {
	if schedule == nil {
		return fmt.Errorf("parameter schedule is required, but nil was provided")
	}

	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_InspectorStartAssessmentRun) validateBindParameters(schedule awscdkscheduleralpha.ISchedule) error {
	if schedule == nil {
		return fmt.Errorf("parameter schedule is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_InspectorStartAssessmentRun) validateBindBaseTargetConfigParameters(_schedule awscdkscheduleralpha.ISchedule) error {
	if _schedule == nil {
		return fmt.Errorf("parameter _schedule is required, but nil was provided")
	}

	return nil
}

func validateNewInspectorStartAssessmentRunParameters(template awsinspector.CfnAssessmentTemplate, props *ScheduleTargetBaseProps) error {
	if template == nil {
		return fmt.Errorf("parameter template is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


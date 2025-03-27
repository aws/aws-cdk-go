//go:build !no_runtime_type_checking

package awscdkscheduleralpha

import (
	"fmt"
)

func (s *jsiiProxy_ScheduleTargetInput) validateBindParameters(schedule ISchedule) error {
	if schedule == nil {
		return fmt.Errorf("parameter schedule is required, but nil was provided")
	}

	return nil
}

func validateScheduleTargetInput_FromObjectParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func validateScheduleTargetInput_FromTextParameters(text *string) error {
	if text == nil {
		return fmt.Errorf("parameter text is required, but nil was provided")
	}

	return nil
}


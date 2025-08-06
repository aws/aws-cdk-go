//go:build !no_runtime_type_checking

package awsscheduler

import (
	"fmt"
)

func (i *jsiiProxy_IScheduleTarget) validateBindParameters(_schedule ISchedule) error {
	if _schedule == nil {
		return fmt.Errorf("parameter _schedule is required, but nil was provided")
	}

	return nil
}


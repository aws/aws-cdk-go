//go:build no_runtime_type_checking

package awsscheduler

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IScheduleTarget) validateBindParameters(_schedule ISchedule) error {
	return nil
}


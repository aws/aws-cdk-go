//go:build no_runtime_type_checking

package awsscheduler

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScheduleTargetInput) validateBindParameters(schedule ISchedule) error {
	return nil
}

func validateScheduleTargetInput_FromObjectParameters(obj interface{}) error {
	return nil
}

func validateScheduleTargetInput_FromTextParameters(text *string) error {
	return nil
}


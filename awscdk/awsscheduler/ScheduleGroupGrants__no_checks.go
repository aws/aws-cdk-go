//go:build no_runtime_type_checking

package awsscheduler

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScheduleGroupGrants) validateDeleteSchedulesParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_ScheduleGroupGrants) validateReadSchedulesParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_ScheduleGroupGrants) validateWriteSchedulesParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateScheduleGroupGrants_FromScheduleGroupParameters(resource interfacesawsscheduler.IScheduleGroupRef) error {
	return nil
}


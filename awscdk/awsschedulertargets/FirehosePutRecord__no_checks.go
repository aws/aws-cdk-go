//go:build no_runtime_type_checking

package awsschedulertargets

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirehosePutRecord) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (f *jsiiProxy_FirehosePutRecord) validateBindParameters(schedule awsscheduler.ISchedule) error {
	return nil
}

func (f *jsiiProxy_FirehosePutRecord) validateBindBaseTargetConfigParameters(_schedule awsscheduler.ISchedule) error {
	return nil
}

func validateNewFirehosePutRecordParameters(deliveryStream awskinesisfirehose.IDeliveryStream, props *ScheduleTargetBaseProps) error {
	return nil
}


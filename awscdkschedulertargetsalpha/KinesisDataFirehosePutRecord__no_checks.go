//go:build no_runtime_type_checking

package awscdkschedulertargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisDataFirehosePutRecord) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (k *jsiiProxy_KinesisDataFirehosePutRecord) validateBindParameters(schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func (k *jsiiProxy_KinesisDataFirehosePutRecord) validateBindBaseTargetConfigParameters(_schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func validateNewKinesisDataFirehosePutRecordParameters(deliveryStream awskinesisfirehose.IDeliveryStream, props *ScheduleTargetBaseProps) error {
	return nil
}


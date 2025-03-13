//go:build no_runtime_type_checking

package awscdkschedulertargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisStreamPutRecord) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (k *jsiiProxy_KinesisStreamPutRecord) validateBindParameters(schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func (k *jsiiProxy_KinesisStreamPutRecord) validateBindBaseTargetConfigParameters(_schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func validateNewKinesisStreamPutRecordParameters(stream awskinesis.IStream, props *KinesisStreamPutRecordProps) error {
	return nil
}


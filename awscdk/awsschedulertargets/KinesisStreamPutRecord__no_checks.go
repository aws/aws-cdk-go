//go:build no_runtime_type_checking

package awsschedulertargets

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisStreamPutRecord) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (k *jsiiProxy_KinesisStreamPutRecord) validateBindParameters(schedule awsscheduler.ISchedule) error {
	return nil
}

func (k *jsiiProxy_KinesisStreamPutRecord) validateBindBaseTargetConfigParameters(_schedule awsscheduler.ISchedule) error {
	return nil
}

func validateNewKinesisStreamPutRecordParameters(stream awskinesis.IStream, props *KinesisStreamPutRecordProps) error {
	return nil
}


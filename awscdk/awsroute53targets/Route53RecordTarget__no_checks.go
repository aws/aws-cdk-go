//go:build no_runtime_type_checking

package awsroute53targets

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Route53RecordTarget) validateBindParameters(_record awsroute53.IRecordSet) error {
	return nil
}

func validateNewRoute53RecordTargetParameters(record awsroute53.IRecordSet) error {
	return nil
}


//go:build no_runtime_type_checking

package awssqs

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueueGrants) validateConsumeMessagesParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (q *jsiiProxy_QueueGrants) validatePurgeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (q *jsiiProxy_QueueGrants) validateSendMessagesParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateQueueGrants_FromQueueParameters(resource interfacesawssqs.IQueueRef) error {
	return nil
}


//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchJob) validateBindParameters(rule awsevents.IRule) error {
	return nil
}

func validateNewBatchJobParameters(jobQueueArn *string, jobQueueScope constructs.IConstruct, jobDefinitionArn *string, jobDefinitionScope constructs.IConstruct, props *BatchJobProps) error {
	return nil
}


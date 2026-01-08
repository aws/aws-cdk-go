//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobQueue) validateAddComputeEnvironmentParameters(computeEnvironment interfacesawsbatch.IComputeEnvironmentRef, order *float64) error {
	return nil
}

func (j *jsiiProxy_JobQueue) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (j *jsiiProxy_JobQueue) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (j *jsiiProxy_JobQueue) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateJobQueue_FromJobQueueArnParameters(scope constructs.Construct, id *string, jobQueueArn *string) error {
	return nil
}

func validateJobQueue_IsConstructParameters(x interface{}) error {
	return nil
}

func validateJobQueue_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateJobQueue_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewJobQueueParameters(scope constructs.Construct, id *string, props *JobQueueProps) error {
	return nil
}


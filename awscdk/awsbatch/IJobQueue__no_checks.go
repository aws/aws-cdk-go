//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IJobQueue) validateAddComputeEnvironmentParameters(computeEnvironment interfacesawsbatch.IComputeEnvironmentRef, order *float64) error {
	return nil
}

func (i *jsiiProxy_IJobQueue) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}


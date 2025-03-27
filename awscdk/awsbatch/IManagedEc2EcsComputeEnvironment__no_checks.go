//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IManagedEc2EcsComputeEnvironment) validateAddInstanceClassParameters(instanceClass awsec2.InstanceClass) error {
	return nil
}

func (i *jsiiProxy_IManagedEc2EcsComputeEnvironment) validateAddInstanceTypeParameters(instanceType awsec2.InstanceType) error {
	return nil
}


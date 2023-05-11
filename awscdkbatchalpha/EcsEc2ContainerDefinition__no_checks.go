//go:build no_runtime_type_checking

// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsEc2ContainerDefinition) validateAddUlimitParameters(ulimit *Ulimit) error {
	return nil
}

func (e *jsiiProxy_EcsEc2ContainerDefinition) validateAddVolumeParameters(volume EcsVolume) error {
	return nil
}

func validateEcsEc2ContainerDefinition_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEcsEc2ContainerDefinitionParameters(scope constructs.Construct, id *string, props *EcsEc2ContainerDefinitionProps) error {
	return nil
}


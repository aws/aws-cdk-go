//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RunEcsEc2Task) validateBindParameters(task awsstepfunctions.Task) error {
	return nil
}

func (r *jsiiProxy_RunEcsEc2Task) validateConfigureAwsVpcNetworkingParameters(vpc awsec2.IVpc, subnetSelection *awsec2.SubnetSelection) error {
	return nil
}

func validateNewRunEcsEc2TaskParameters(props *RunEcsEc2TaskProps) error {
	return nil
}


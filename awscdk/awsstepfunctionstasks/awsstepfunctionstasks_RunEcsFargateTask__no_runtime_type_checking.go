//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RunEcsFargateTask) validateBindParameters(task awsstepfunctions.Task) error {
	return nil
}

func (r *jsiiProxy_RunEcsFargateTask) validateConfigureAwsVpcNetworkingParameters(vpc awsec2.IVpc, subnetSelection *awsec2.SubnetSelection) error {
	return nil
}

func validateNewRunEcsFargateTaskParameters(props *RunEcsFargateTaskProps) error {
	return nil
}


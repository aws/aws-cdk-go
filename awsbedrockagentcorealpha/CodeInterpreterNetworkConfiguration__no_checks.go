//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateCodeInterpreterNetworkConfiguration_UsingVpcParameters(scope constructs.Construct, vpcConfig *VpcConfigProps) error {
	return nil
}

func validateNewCodeInterpreterNetworkConfigurationParameters(mode *string, vpcConfig *VpcConfigProps) error {
	return nil
}


//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCodeDeployBlueGreenHook) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateCfnCodeDeployBlueGreenHook_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnCodeDeployBlueGreenHook_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) validateSetAdditionalOptionsParameters(val *CfnCodeDeployBlueGreenAdditionalOptions) error {
	return nil
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) validateSetApplicationsParameters(val *[]*CfnCodeDeployBlueGreenApplication) error {
	return nil
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) validateSetLifecycleEventHooksParameters(val *CfnCodeDeployBlueGreenLifecycleEventHooks) error {
	return nil
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) validateSetServiceRoleParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CfnCodeDeployBlueGreenHook) validateSetTrafficRoutingConfigParameters(val *CfnTrafficRoutingConfig) error {
	return nil
}

func validateNewCfnCodeDeployBlueGreenHookParameters(scope constructs.Construct, id *string, props *CfnCodeDeployBlueGreenHookProps) error {
	return nil
}


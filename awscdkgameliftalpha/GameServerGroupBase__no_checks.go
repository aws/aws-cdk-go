//go:build no_runtime_type_checking

package awscdkgameliftalpha

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GameServerGroupBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_GameServerGroupBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GameServerGroupBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (g *jsiiProxy_GameServerGroupBase) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (g *jsiiProxy_GameServerGroupBase) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateGameServerGroupBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGameServerGroupBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGameServerGroupBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGameServerGroupBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}


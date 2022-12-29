//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Dashboard) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_Dashboard) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_Dashboard) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateDashboard_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDashboard_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDashboard_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDashboardParameters(scope constructs.Construct, id *string, props *DashboardProps) error {
	return nil
}


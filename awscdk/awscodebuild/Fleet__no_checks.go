//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_Fleet) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_Fleet) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_Fleet) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateFleet_FromFleetArnParameters(scope constructs.Construct, id *string, fleetArn *string) error {
	return nil
}

func validateFleet_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFleet_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFleet_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFleetParameters(scope constructs.Construct, id *string, props *FleetProps) error {
	return nil
}


//go:build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DedicatedIpPool) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DedicatedIpPool) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DedicatedIpPool) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateDedicatedIpPool_FromDedicatedIpPoolNameParameters(scope constructs.Construct, id *string, dedicatedIpPoolName *string) error {
	return nil
}

func validateDedicatedIpPool_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDedicatedIpPool_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDedicatedIpPool_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDedicatedIpPoolParameters(scope constructs.Construct, id *string, props *DedicatedIpPoolProps) error {
	return nil
}


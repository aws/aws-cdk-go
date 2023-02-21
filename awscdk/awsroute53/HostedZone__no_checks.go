//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HostedZone) validateAddVpcParameters(vpc awsec2.IVpc) error {
	return nil
}

func (h *jsiiProxy_HostedZone) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (h *jsiiProxy_HostedZone) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (h *jsiiProxy_HostedZone) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateHostedZone_FromHostedZoneAttributesParameters(scope constructs.Construct, id *string, attrs *HostedZoneAttributes) error {
	return nil
}

func validateHostedZone_FromHostedZoneIdParameters(scope constructs.Construct, id *string, hostedZoneId *string) error {
	return nil
}

func validateHostedZone_FromLookupParameters(scope constructs.Construct, id *string, query *HostedZoneProviderProps) error {
	return nil
}

func validateHostedZone_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHostedZone_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateHostedZone_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewHostedZoneParameters(scope constructs.Construct, id *string, props *HostedZoneProps) error {
	return nil
}


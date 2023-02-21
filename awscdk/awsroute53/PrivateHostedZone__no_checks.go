//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateHostedZone) validateAddVpcParameters(vpc awsec2.IVpc) error {
	return nil
}

func (p *jsiiProxy_PrivateHostedZone) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PrivateHostedZone) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PrivateHostedZone) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePrivateHostedZone_FromHostedZoneAttributesParameters(scope constructs.Construct, id *string, attrs *HostedZoneAttributes) error {
	return nil
}

func validatePrivateHostedZone_FromHostedZoneIdParameters(scope constructs.Construct, id *string, hostedZoneId *string) error {
	return nil
}

func validatePrivateHostedZone_FromLookupParameters(scope constructs.Construct, id *string, query *HostedZoneProviderProps) error {
	return nil
}

func validatePrivateHostedZone_FromPrivateHostedZoneIdParameters(scope constructs.Construct, id *string, privateHostedZoneId *string) error {
	return nil
}

func validatePrivateHostedZone_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePrivateHostedZone_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePrivateHostedZone_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPrivateHostedZoneParameters(scope constructs.Construct, id *string, props *PrivateHostedZoneProps) error {
	return nil
}


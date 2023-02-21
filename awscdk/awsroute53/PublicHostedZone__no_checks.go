//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PublicHostedZone) validateAddDelegationParameters(delegate IPublicHostedZone, opts *ZoneDelegationOptions) error {
	return nil
}

func (p *jsiiProxy_PublicHostedZone) validateAddVpcParameters(_vpc awsec2.IVpc) error {
	return nil
}

func (p *jsiiProxy_PublicHostedZone) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PublicHostedZone) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PublicHostedZone) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_PublicHostedZone) validateGrantDelegationParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validatePublicHostedZone_FromHostedZoneAttributesParameters(scope constructs.Construct, id *string, attrs *HostedZoneAttributes) error {
	return nil
}

func validatePublicHostedZone_FromHostedZoneIdParameters(scope constructs.Construct, id *string, hostedZoneId *string) error {
	return nil
}

func validatePublicHostedZone_FromLookupParameters(scope constructs.Construct, id *string, query *HostedZoneProviderProps) error {
	return nil
}

func validatePublicHostedZone_FromPublicHostedZoneAttributesParameters(scope constructs.Construct, id *string, attrs *PublicHostedZoneAttributes) error {
	return nil
}

func validatePublicHostedZone_FromPublicHostedZoneIdParameters(scope constructs.Construct, id *string, publicHostedZoneId *string) error {
	return nil
}

func validatePublicHostedZone_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePublicHostedZone_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePublicHostedZone_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPublicHostedZoneParameters(scope constructs.Construct, id *string, props *PublicHostedZoneProps) error {
	return nil
}


//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HostedZoneGrants) validateDelegationParameters(grantee awsiam.IGrantable, delegationOptions *GrantDelegationOptions) error {
	return nil
}

func validateHostedZoneGrants_FromHostedZoneParameters(hostedZone INamedHostedZoneRef) error {
	return nil
}


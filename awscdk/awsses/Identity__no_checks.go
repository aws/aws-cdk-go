//go:build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func validateIdentity_DomainParameters(domain *string) error {
	return nil
}

func validateIdentity_EmailParameters(email *string) error {
	return nil
}

func validateIdentity_PublicHostedZoneParameters(hostedZone awsroute53.IPublicHostedZone) error {
	return nil
}


//go:build no_runtime_type_checking

package awscdkroute53resolveralpha

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirewallDomains) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateFirewallDomains_FromAssetParameters(assetPath *string) error {
	return nil
}

func validateFirewallDomains_FromListParameters(list *[]*string) error {
	return nil
}

func validateFirewallDomains_FromS3Parameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateFirewallDomains_FromS3UrlParameters(url *string) error {
	return nil
}


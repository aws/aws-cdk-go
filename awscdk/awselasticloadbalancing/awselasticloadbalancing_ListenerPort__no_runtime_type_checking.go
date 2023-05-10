//go:build no_runtime_type_checking

package awselasticloadbalancing

// Building without runtime type checking enabled, so all the below just return nil

func validateNewListenerPortParameters(securityGroup awsec2.ISecurityGroup, defaultPort awsec2.Port) error {
	return nil
}


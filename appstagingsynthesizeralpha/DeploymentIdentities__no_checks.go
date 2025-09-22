//go:build no_runtime_type_checking

package appstagingsynthesizeralpha

// Building without runtime type checking enabled, so all the below just return nil

func validateDeploymentIdentities_DefaultBootstrapRolesParameters(options *DefaultBootstrapRolesOptions) error {
	return nil
}

func validateDeploymentIdentities_SpecifyRolesParameters(roles *BootstrapRoles) error {
	return nil
}


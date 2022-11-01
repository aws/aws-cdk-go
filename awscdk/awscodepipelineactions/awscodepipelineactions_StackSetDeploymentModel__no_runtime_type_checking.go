//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func validateStackSetDeploymentModel_OrganizationsParameters(props *OrganizationsDeploymentProps) error {
	return nil
}

func validateStackSetDeploymentModel_SelfManagedParameters(props *SelfManagedDeploymentProps) error {
	return nil
}


//go:build no_runtime_type_checking

package awscodecommit

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RepositoryGrants) validatePullParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RepositoryGrants) validatePullPushParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RepositoryGrants) validateReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateRepositoryGrants_FromRepositoryParameters(resource interfacesawscodecommit.IRepositoryRef) error {
	return nil
}


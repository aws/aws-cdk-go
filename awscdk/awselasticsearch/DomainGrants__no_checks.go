//go:build no_runtime_type_checking

package awselasticsearch

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DomainGrants) validateReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DomainGrants) validateReadWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DomainGrants) validateWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateDomainGrants_FromDomainParameters(resource interfacesawselasticsearch.IDomainRef) error {
	return nil
}


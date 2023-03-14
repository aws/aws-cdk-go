//go:build no_runtime_type_checking

package awsroute53targets

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserPoolDomainTarget) validateBindParameters(_record awsroute53.IRecordSet) error {
	return nil
}

func validateNewUserPoolDomainTargetParameters(domain awscognito.UserPoolDomain) error {
	return nil
}


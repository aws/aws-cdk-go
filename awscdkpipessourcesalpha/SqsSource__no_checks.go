//go:build no_runtime_type_checking

package awscdkpipessourcesalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqsSource) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (s *jsiiProxy_SqsSource) validateGrantReadParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewSqsSourceParameters(queue awssqs.IQueue, parameters *SqsSourceParameters) error {
	return nil
}


//go:build no_runtime_type_checking

package awsroute53targets

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ElasticBeanstalkEnvironmentEndpointTarget) validateBindParameters(_record awsroute53.IRecordSet) error {
	return nil
}

func validateNewElasticBeanstalkEnvironmentEndpointTargetParameters(environmentEndpoint *string) error {
	return nil
}


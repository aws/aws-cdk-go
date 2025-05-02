//go:build no_runtime_type_checking

package awscdkpipestargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqsTarget) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (s *jsiiProxy_SqsTarget) validateGrantPushParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewSqsTargetParameters(queue awssqs.IQueue, parameters *SqsTargetParameters) error {
	return nil
}


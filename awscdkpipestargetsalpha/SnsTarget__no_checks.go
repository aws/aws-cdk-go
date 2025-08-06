//go:build no_runtime_type_checking

package awscdkpipestargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnsTarget) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (s *jsiiProxy_SnsTarget) validateGrantPushParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewSnsTargetParameters(topic awssns.ITopic, parameters *SnsTargetParameters) error {
	return nil
}


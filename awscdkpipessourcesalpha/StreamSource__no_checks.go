//go:build no_runtime_type_checking

package awscdkpipessourcesalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StreamSource) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (s *jsiiProxy_StreamSource) validateGetDeadLetterTargetArnParameters(deadLetterTarget interface{}) error {
	return nil
}

func (s *jsiiProxy_StreamSource) validateGrantPushParameters(grantee awsiam.IRole, deadLetterTarget interface{}) error {
	return nil
}

func (s *jsiiProxy_StreamSource) validateGrantReadParameters(grantee awsiam.IRole) error {
	return nil
}

func validateStreamSource_IsSourceWithDeadLetterTargetParameters(source awscdkpipesalpha.ISource) error {
	return nil
}

func validateNewStreamSourceParameters(sourceArn *string, sourceParameters *StreamSourceParameters) error {
	return nil
}


//go:build no_runtime_type_checking

package awscdkpipesalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SourceWithDeadLetterTarget) validateBindParameters(pipe IPipe) error {
	return nil
}

func (s *jsiiProxy_SourceWithDeadLetterTarget) validateGetDeadLetterTargetArnParameters(deadLetterTarget interface{}) error {
	return nil
}

func (s *jsiiProxy_SourceWithDeadLetterTarget) validateGrantPushParameters(grantee awsiam.IRole, deadLetterTarget interface{}) error {
	return nil
}

func (s *jsiiProxy_SourceWithDeadLetterTarget) validateGrantReadParameters(grantee awsiam.IRole) error {
	return nil
}

func validateSourceWithDeadLetterTarget_IsSourceWithDeadLetterTargetParameters(source ISource) error {
	return nil
}

func validateNewSourceWithDeadLetterTargetParameters(sourceArn *string, deadLetterTarget interface{}) error {
	return nil
}


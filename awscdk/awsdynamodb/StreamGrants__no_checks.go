//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StreamGrants) validateActionsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StreamGrants) validateListParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StreamGrants) validateReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateNewStreamGrantsParameters(props *StreamGrantsProps) error {
	return nil
}


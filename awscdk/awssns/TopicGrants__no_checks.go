//go:build no_runtime_type_checking

package awssns

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TopicGrants) validatePublishParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TopicGrants) validateSubscribeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateTopicGrants_FromTopicParameters(resource interfacesawssns.ITopicRef) error {
	return nil
}


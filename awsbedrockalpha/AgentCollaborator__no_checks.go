//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AgentCollaborator) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateNewAgentCollaboratorParameters(props *AgentCollaboratorProps) error {
	return nil
}


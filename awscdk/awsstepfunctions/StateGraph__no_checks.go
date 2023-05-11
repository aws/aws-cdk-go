//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StateGraph) validateRegisterPolicyStatementParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_StateGraph) validateRegisterStateParameters(state State) error {
	return nil
}

func (s *jsiiProxy_StateGraph) validateRegisterSuperGraphParameters(graph StateGraph) error {
	return nil
}

func validateNewStateGraphParameters(startState State, graphDescription *string) error {
	return nil
}


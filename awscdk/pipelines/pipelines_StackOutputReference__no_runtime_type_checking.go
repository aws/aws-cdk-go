//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StackOutputReference) validateIsProducedByParameters(stack StackDeployment) error {
	return nil
}

func validateStackOutputReference_FromCfnOutputParameters(output awscdk.CfnOutput) error {
	return nil
}


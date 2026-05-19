//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func validateAgentCoreRuntime_OfParameters(value *string) error {
	return nil
}


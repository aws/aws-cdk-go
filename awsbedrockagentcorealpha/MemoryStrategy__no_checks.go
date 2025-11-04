//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateMemoryStrategy_UsingSelfManagedParameters(config *SelfManagedStrategyProps) error {
	return nil
}

func validateMemoryStrategy_UsingSemanticParameters(config *ManagedStrategyProps) error {
	return nil
}

func validateMemoryStrategy_UsingSummarizationParameters(config *ManagedStrategyProps) error {
	return nil
}

func validateMemoryStrategy_UsingUserPreferenceParameters(config *ManagedStrategyProps) error {
	return nil
}


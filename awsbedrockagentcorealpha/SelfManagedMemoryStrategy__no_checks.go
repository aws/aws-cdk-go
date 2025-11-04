//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SelfManagedMemoryStrategy) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateNewSelfManagedMemoryStrategyParameters(strategyType MemoryStrategyType, props *SelfManagedStrategyProps) error {
	return nil
}


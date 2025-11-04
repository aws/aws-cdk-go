//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedMemoryStrategy) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateNewManagedMemoryStrategyParameters(strategyType MemoryStrategyType, props *ManagedStrategyProps) error {
	return nil
}


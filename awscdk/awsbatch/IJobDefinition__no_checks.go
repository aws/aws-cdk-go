//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IJobDefinition) validateAddRetryStrategyParameters(strategy RetryStrategy) error {
	return nil
}


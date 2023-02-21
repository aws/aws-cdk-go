//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func validateLogRetention_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLogRetentionParameters(scope constructs.Construct, id *string, props *LogRetentionProps) error {
	return nil
}


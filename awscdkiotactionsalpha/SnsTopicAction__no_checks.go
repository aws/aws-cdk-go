//go:build no_runtime_type_checking

package awscdkiotactionsalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateNewSnsTopicActionParameters(topic awssns.ITopic, props *SnsTopicActionProps) error {
	return nil
}


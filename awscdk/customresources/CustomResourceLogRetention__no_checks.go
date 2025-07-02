//go:build no_runtime_type_checking

package customresources

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomResourceLogRetention) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func validateNewCustomResourceLogRetentionParameters(setLogRetention awslogs.RetentionDays) error {
	return nil
}


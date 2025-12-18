//go:build no_runtime_type_checking

package previewawslogs

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ILogsDelivery) validateBindParameters(scope constructs.IConstruct, logType *string, sourceResourceArn *string) error {
	return nil
}


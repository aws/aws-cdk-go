//go:build no_runtime_type_checking

package previewawslogs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3LogsDelivery) validateBindParameters(scope constructs.IConstruct, logType *string, sourceResourceArn *string) error {
	return nil
}

func validateNewS3LogsDeliveryParameters(bucket interfacesawss3.IBucketRef, props *S3LogsDeliveryProps) error {
	return nil
}


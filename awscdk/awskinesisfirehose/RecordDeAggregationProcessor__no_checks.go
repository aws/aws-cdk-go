//go:build no_runtime_type_checking

package awskinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RecordDeAggregationProcessor) validateBindParameters(scope constructs.Construct, options *DataProcessorBindOptions) error {
	return nil
}

func validateRecordDeAggregationProcessor_DelimitedParameters(delimiter *string) error {
	return nil
}

func validateNewRecordDeAggregationProcessorParameters(options *RecordDeAggregationProcessorOptions) error {
	return nil
}


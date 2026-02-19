//go:build no_runtime_type_checking

package awskinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MetadataExtractionProcessor) validateBindParameters(scope constructs.Construct, options *DataProcessorBindOptions) error {
	return nil
}

func validateMetadataExtractionProcessor_Jq16Parameters(query *map[string]*string) error {
	return nil
}

func validateNewMetadataExtractionProcessorParameters(options *MetadataExtractionProcessorOptions, keys *[]*string) error {
	return nil
}


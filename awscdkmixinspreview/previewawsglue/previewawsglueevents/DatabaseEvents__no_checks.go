//go:build no_runtime_type_checking

package previewawsglueevents

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseEvents) validateGlueDataCatalogDatabaseStateChangePatternParameters(options *DatabaseEvents_GlueDataCatalogDatabaseStateChange_GlueDataCatalogDatabaseStateChangeProps) error {
	return nil
}

func (d *jsiiProxy_DatabaseEvents) validateGlueDataCatalogTableStateChangePatternParameters(options *DatabaseEvents_GlueDataCatalogTableStateChange_GlueDataCatalogTableStateChangeProps) error {
	return nil
}

func validateDatabaseEvents_FromDatabaseParameters(databaseRef interfacesawsglue.IDatabaseRef) error {
	return nil
}


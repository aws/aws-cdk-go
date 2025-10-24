//go:build no_runtime_type_checking

package awskinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SchemaConfiguration) validateBindParameters(scope constructs.Construct, options *SchemaConfigurationBindOptions) error {
	return nil
}

func validateSchemaConfiguration_FromCfnTableParameters(table awsglue.CfnTable, props *SchemaConfigurationFromCfnTableProps) error {
	return nil
}


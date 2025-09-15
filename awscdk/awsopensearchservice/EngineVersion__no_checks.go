//go:build no_runtime_type_checking

package awsopensearchservice

// Building without runtime type checking enabled, so all the below just return nil

func validateEngineVersion_ElasticsearchParameters(version *string) error {
	return nil
}

func validateEngineVersion_OpenSearchParameters(version *string) error {
	return nil
}


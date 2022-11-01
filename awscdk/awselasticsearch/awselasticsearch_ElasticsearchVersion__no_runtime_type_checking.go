//go:build no_runtime_type_checking

package awselasticsearch

// Building without runtime type checking enabled, so all the below just return nil

func validateElasticsearchVersion_OfParameters(version *string) error {
	return nil
}


//go:build no_runtime_type_checking

package awsglue

// Building without runtime type checking enabled, so all the below just return nil

func validateWorkerType_OfParameters(workerType *string) error {
	return nil
}


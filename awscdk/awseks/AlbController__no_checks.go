//go:build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func validateAlbController_CreateParameters(scope constructs.Construct, props *AlbControllerProps) error {
	return nil
}

func validateAlbController_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAlbControllerParameters(scope constructs.Construct, id *string, props *AlbControllerProps) error {
	return nil
}


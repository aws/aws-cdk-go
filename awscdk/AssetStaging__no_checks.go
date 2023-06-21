//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetStaging) validateRelativeStagedPathParameters(stack Stack) error {
	return nil
}

func validateAssetStaging_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAssetStagingParameters(scope constructs.Construct, id *string, props *AssetStagingProps) error {
	return nil
}


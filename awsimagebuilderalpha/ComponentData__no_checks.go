//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateComponentData_FromAssetParameters(scope constructs.Construct, id *string, path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateComponentData_FromComponentDocumentJsonObjectParameters(data *ComponentDocument) error {
	return nil
}

func validateComponentData_FromInlineParameters(data *string) error {
	return nil
}

func validateComponentData_FromJsonObjectParameters(data *map[string]interface{}) error {
	return nil
}

func validateComponentData_FromS3Parameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateNewComponentDataParameters(value *string) error {
	return nil
}


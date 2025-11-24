//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateComponentParameterValue_FromStringParameters(value *string) error {
	return nil
}

func validateNewComponentParameterValueParameters(value *[]*string) error {
	return nil
}


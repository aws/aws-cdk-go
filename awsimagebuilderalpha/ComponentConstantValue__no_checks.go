//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateComponentConstantValue_FromStringParameters(value *string) error {
	return nil
}

func validateNewComponentConstantValueParameters(type_ *string, value interface{}) error {
	return nil
}


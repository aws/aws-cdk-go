//go:build no_runtime_type_checking

// The CDK Construct Library for AWS::Glue
package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateSchema_ArrayParameters(itemType *Type) error {
	return nil
}

func validateSchema_CharParameters(length *float64) error {
	return nil
}

func validateSchema_DecimalParameters(precision *float64) error {
	return nil
}

func validateSchema_MapParameters(keyType *Type, valueType *Type) error {
	return nil
}

func validateSchema_StructParameters(columns *[]*Column) error {
	return nil
}

func validateSchema_VarcharParameters(length *float64) error {
	return nil
}


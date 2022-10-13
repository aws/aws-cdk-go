//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsglue

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateSchema_ArrayParameters(itemType *Type) error {
	if itemType == nil {
		return fmt.Errorf("parameter itemType is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(itemType, func() string { return "parameter itemType" }); err != nil {
		return err
	}

	return nil
}

func validateSchema_CharParameters(length *float64) error {
	if length == nil {
		return fmt.Errorf("parameter length is required, but nil was provided")
	}

	return nil
}

func validateSchema_DecimalParameters(precision *float64) error {
	if precision == nil {
		return fmt.Errorf("parameter precision is required, but nil was provided")
	}

	return nil
}

func validateSchema_MapParameters(keyType *Type, valueType *Type) error {
	if keyType == nil {
		return fmt.Errorf("parameter keyType is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(keyType, func() string { return "parameter keyType" }); err != nil {
		return err
	}

	if valueType == nil {
		return fmt.Errorf("parameter valueType is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(valueType, func() string { return "parameter valueType" }); err != nil {
		return err
	}

	return nil
}

func validateSchema_StructParameters(columns *[]*Column) error {
	if columns == nil {
		return fmt.Errorf("parameter columns is required, but nil was provided")
	}
	for idx_0e624a, v := range *columns {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter columns[%#v]", idx_0e624a) }); err != nil {
			return err
		}
	}

	return nil
}

func validateSchema_VarcharParameters(length *float64) error {
	if length == nil {
		return fmt.Errorf("parameter length is required, but nil was provided")
	}

	return nil
}


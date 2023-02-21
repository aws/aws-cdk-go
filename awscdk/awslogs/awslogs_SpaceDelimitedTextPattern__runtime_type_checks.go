//go:build !no_runtime_type_checking

package awslogs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (s *jsiiProxy_SpaceDelimitedTextPattern) validateWhereNumberParameters(columnName *string, comparison *string, value *float64) error {
	if columnName == nil {
		return fmt.Errorf("parameter columnName is required, but nil was provided")
	}

	if comparison == nil {
		return fmt.Errorf("parameter comparison is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SpaceDelimitedTextPattern) validateWhereStringParameters(columnName *string, comparison *string, value *string) error {
	if columnName == nil {
		return fmt.Errorf("parameter columnName is required, but nil was provided")
	}

	if comparison == nil {
		return fmt.Errorf("parameter comparison is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateSpaceDelimitedTextPattern_ConstructParameters(columns *[]*string) error {
	if columns == nil {
		return fmt.Errorf("parameter columns is required, but nil was provided")
	}

	return nil
}

func validateNewSpaceDelimitedTextPatternParameters(columns *[]*string, restrictions *map[string]*[]*ColumnRestriction) error {
	if columns == nil {
		return fmt.Errorf("parameter columns is required, but nil was provided")
	}

	if restrictions == nil {
		return fmt.Errorf("parameter restrictions is required, but nil was provided")
	}
	for idx_15a956, v := range *restrictions {
		for idx_4c9448, v := range *v {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter restrictions[%#v][%#v]", idx_15a956, idx_4c9448) }); err != nil {
				return err
			}
		}
	}

	return nil
}


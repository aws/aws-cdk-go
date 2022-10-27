//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SpaceDelimitedTextPattern) validateWhereNumberParameters(columnName *string, comparison *string, value *float64) error {
	return nil
}

func (s *jsiiProxy_SpaceDelimitedTextPattern) validateWhereStringParameters(columnName *string, comparison *string, value *string) error {
	return nil
}

func validateSpaceDelimitedTextPattern_ConstructParameters(columns *[]*string) error {
	return nil
}

func validateNewSpaceDelimitedTextPatternParameters(columns *[]*string, restrictions *map[string]*[]*ColumnRestriction) error {
	return nil
}


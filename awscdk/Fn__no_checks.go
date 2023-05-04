//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateFn_Base64Parameters(data *string) error {
	return nil
}

func validateFn_CidrParameters(ipBlock *string, count *float64) error {
	return nil
}

func validateFn_ConditionContainsParameters(listOfStrings *[]*string, value *string) error {
	return nil
}

func validateFn_ConditionEachMemberEqualsParameters(listOfStrings *[]*string, value *string) error {
	return nil
}

func validateFn_ConditionEachMemberInParameters(stringsToCheck *[]*string, stringsToMatch *[]*string) error {
	return nil
}

func validateFn_ConditionEqualsParameters(lhs interface{}, rhs interface{}) error {
	return nil
}

func validateFn_ConditionIfParameters(conditionId *string, valueIfTrue interface{}, valueIfFalse interface{}) error {
	return nil
}

func validateFn_ConditionNotParameters(condition ICfnConditionExpression) error {
	return nil
}

func validateFn_FindInMapParameters(mapName *string, topLevelKey *string, secondLevelKey *string) error {
	return nil
}

func validateFn_GetAttParameters(logicalNameOfResource *string, attributeName *string) error {
	return nil
}

func validateFn_ImportListValueParameters(sharedValueToImport *string, assumedLength *float64) error {
	return nil
}

func validateFn_ImportValueParameters(sharedValueToImport *string) error {
	return nil
}

func validateFn_JoinParameters(delimiter *string, listOfValues *[]*string) error {
	return nil
}

func validateFn_LenParameters(array interface{}) error {
	return nil
}

func validateFn_ParseDomainNameParameters(url *string) error {
	return nil
}

func validateFn_RefParameters(logicalName *string) error {
	return nil
}

func validateFn_RefAllParameters(parameterType *string) error {
	return nil
}

func validateFn_SelectParameters(index *float64, array *[]*string) error {
	return nil
}

func validateFn_SplitParameters(delimiter *string, source *string) error {
	return nil
}

func validateFn_SubParameters(body *string) error {
	return nil
}

func validateFn_ToJsonStringParameters(object interface{}) error {
	return nil
}

func validateFn_TransformParameters(macroName *string, parameters *map[string]interface{}) error {
	return nil
}

func validateFn_ValueOfParameters(parameterOrLogicalId *string, attribute *string) error {
	return nil
}

func validateFn_ValueOfAllParameters(parameterType *string, attribute *string) error {
	return nil
}


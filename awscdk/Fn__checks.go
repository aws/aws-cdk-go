//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"
)

func validateFn_Base64Parameters(data *string) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}

	return nil
}

func validateFn_CidrParameters(ipBlock *string, count *float64) error {
	if ipBlock == nil {
		return fmt.Errorf("parameter ipBlock is required, but nil was provided")
	}

	if count == nil {
		return fmt.Errorf("parameter count is required, but nil was provided")
	}

	return nil
}

func validateFn_ConditionContainsParameters(listOfStrings *[]*string, value *string) error {
	if listOfStrings == nil {
		return fmt.Errorf("parameter listOfStrings is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_ConditionEachMemberEqualsParameters(listOfStrings *[]*string, value *string) error {
	if listOfStrings == nil {
		return fmt.Errorf("parameter listOfStrings is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_ConditionEachMemberInParameters(stringsToCheck *[]*string, stringsToMatch *[]*string) error {
	if stringsToCheck == nil {
		return fmt.Errorf("parameter stringsToCheck is required, but nil was provided")
	}

	if stringsToMatch == nil {
		return fmt.Errorf("parameter stringsToMatch is required, but nil was provided")
	}

	return nil
}

func validateFn_ConditionEqualsParameters(lhs interface{}, rhs interface{}) error {
	if lhs == nil {
		return fmt.Errorf("parameter lhs is required, but nil was provided")
	}

	if rhs == nil {
		return fmt.Errorf("parameter rhs is required, but nil was provided")
	}

	return nil
}

func validateFn_ConditionIfParameters(conditionId *string, valueIfTrue interface{}, valueIfFalse interface{}) error {
	if conditionId == nil {
		return fmt.Errorf("parameter conditionId is required, but nil was provided")
	}

	if valueIfTrue == nil {
		return fmt.Errorf("parameter valueIfTrue is required, but nil was provided")
	}

	if valueIfFalse == nil {
		return fmt.Errorf("parameter valueIfFalse is required, but nil was provided")
	}

	return nil
}

func validateFn_ConditionNotParameters(condition ICfnConditionExpression) error {
	if condition == nil {
		return fmt.Errorf("parameter condition is required, but nil was provided")
	}

	return nil
}

func validateFn_FindInMapParameters(mapName *string, topLevelKey *string, secondLevelKey *string) error {
	if mapName == nil {
		return fmt.Errorf("parameter mapName is required, but nil was provided")
	}

	if topLevelKey == nil {
		return fmt.Errorf("parameter topLevelKey is required, but nil was provided")
	}

	if secondLevelKey == nil {
		return fmt.Errorf("parameter secondLevelKey is required, but nil was provided")
	}

	return nil
}

func validateFn_GetAttParameters(logicalNameOfResource *string, attributeName *string) error {
	if logicalNameOfResource == nil {
		return fmt.Errorf("parameter logicalNameOfResource is required, but nil was provided")
	}

	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func validateFn_ImportListValueParameters(sharedValueToImport *string, assumedLength *float64) error {
	if sharedValueToImport == nil {
		return fmt.Errorf("parameter sharedValueToImport is required, but nil was provided")
	}

	if assumedLength == nil {
		return fmt.Errorf("parameter assumedLength is required, but nil was provided")
	}

	return nil
}

func validateFn_ImportValueParameters(sharedValueToImport *string) error {
	if sharedValueToImport == nil {
		return fmt.Errorf("parameter sharedValueToImport is required, but nil was provided")
	}

	return nil
}

func validateFn_JoinParameters(delimiter *string, listOfValues *[]*string) error {
	if delimiter == nil {
		return fmt.Errorf("parameter delimiter is required, but nil was provided")
	}

	if listOfValues == nil {
		return fmt.Errorf("parameter listOfValues is required, but nil was provided")
	}

	return nil
}

func validateFn_LenParameters(array interface{}) error {
	if array == nil {
		return fmt.Errorf("parameter array is required, but nil was provided")
	}

	return nil
}

func validateFn_ParseDomainNameParameters(url *string) error {
	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	return nil
}

func validateFn_RefParameters(logicalName *string) error {
	if logicalName == nil {
		return fmt.Errorf("parameter logicalName is required, but nil was provided")
	}

	return nil
}

func validateFn_RefAllParameters(parameterType *string) error {
	if parameterType == nil {
		return fmt.Errorf("parameter parameterType is required, but nil was provided")
	}

	return nil
}

func validateFn_SelectParameters(index *float64, array *[]*string) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	if array == nil {
		return fmt.Errorf("parameter array is required, but nil was provided")
	}

	return nil
}

func validateFn_SplitParameters(delimiter *string, source *string) error {
	if delimiter == nil {
		return fmt.Errorf("parameter delimiter is required, but nil was provided")
	}

	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}

	return nil
}

func validateFn_SubParameters(body *string) error {
	if body == nil {
		return fmt.Errorf("parameter body is required, but nil was provided")
	}

	return nil
}

func validateFn_ToJsonStringParameters(object interface{}) error {
	if object == nil {
		return fmt.Errorf("parameter object is required, but nil was provided")
	}

	return nil
}

func validateFn_TransformParameters(macroName *string, parameters *map[string]interface{}) error {
	if macroName == nil {
		return fmt.Errorf("parameter macroName is required, but nil was provided")
	}

	if parameters == nil {
		return fmt.Errorf("parameter parameters is required, but nil was provided")
	}

	return nil
}

func validateFn_ValueOfParameters(parameterOrLogicalId *string, attribute *string) error {
	if parameterOrLogicalId == nil {
		return fmt.Errorf("parameter parameterOrLogicalId is required, but nil was provided")
	}

	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func validateFn_ValueOfAllParameters(parameterType *string, attribute *string) error {
	if parameterType == nil {
		return fmt.Errorf("parameter parameterType is required, but nil was provided")
	}

	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}


//go:build !no_runtime_type_checking

package cloudformationinclude

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CfnInclude) validateGetConditionParameters(conditionName *string) error {
	if conditionName == nil {
		return fmt.Errorf("parameter conditionName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInclude) validateGetHookParameters(hookLogicalId *string) error {
	if hookLogicalId == nil {
		return fmt.Errorf("parameter hookLogicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInclude) validateGetMappingParameters(mappingName *string) error {
	if mappingName == nil {
		return fmt.Errorf("parameter mappingName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInclude) validateGetNestedStackParameters(logicalId *string) error {
	if logicalId == nil {
		return fmt.Errorf("parameter logicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInclude) validateGetOutputParameters(logicalId *string) error {
	if logicalId == nil {
		return fmt.Errorf("parameter logicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInclude) validateGetParameterParameters(parameterName *string) error {
	if parameterName == nil {
		return fmt.Errorf("parameter parameterName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInclude) validateGetResourceParameters(logicalId *string) error {
	if logicalId == nil {
		return fmt.Errorf("parameter logicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInclude) validateGetRuleParameters(ruleName *string) error {
	if ruleName == nil {
		return fmt.Errorf("parameter ruleName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInclude) validateLoadNestedStackParameters(logicalId *string, nestedStackProps *CfnIncludeProps) error {
	if logicalId == nil {
		return fmt.Errorf("parameter logicalId is required, but nil was provided")
	}

	if nestedStackProps == nil {
		return fmt.Errorf("parameter nestedStackProps is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(nestedStackProps, func() string { return "parameter nestedStackProps" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnInclude) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func validateCfnInclude_IsCfnElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnInclude_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewCfnIncludeParameters(scope constructs.Construct, id *string, props *CfnIncludeProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


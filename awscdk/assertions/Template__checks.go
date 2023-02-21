//go:build !no_runtime_type_checking

package assertions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (t *jsiiProxy_Template) validateAllResourcesParameters(type_ *string, props interface{}) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateAllResourcesPropertiesParameters(type_ *string, props interface{}) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateFindConditionsParameters(logicalId *string) error {
	if logicalId == nil {
		return fmt.Errorf("parameter logicalId is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateFindMappingsParameters(logicalId *string) error {
	if logicalId == nil {
		return fmt.Errorf("parameter logicalId is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateFindOutputsParameters(logicalId *string) error {
	if logicalId == nil {
		return fmt.Errorf("parameter logicalId is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateFindParametersParameters(logicalId *string) error {
	if logicalId == nil {
		return fmt.Errorf("parameter logicalId is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateFindResourcesParameters(type_ *string) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateHasConditionParameters(logicalId *string, props interface{}) error {
	if logicalId == nil {
		return fmt.Errorf("parameter logicalId is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateHasMappingParameters(logicalId *string, props interface{}) error {
	if logicalId == nil {
		return fmt.Errorf("parameter logicalId is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateHasOutputParameters(logicalId *string, props interface{}) error {
	if logicalId == nil {
		return fmt.Errorf("parameter logicalId is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateHasParameterParameters(logicalId *string, props interface{}) error {
	if logicalId == nil {
		return fmt.Errorf("parameter logicalId is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateHasResourceParameters(type_ *string, props interface{}) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateHasResourcePropertiesParameters(type_ *string, props interface{}) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateResourceCountIsParameters(type_ *string, count *float64) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if count == nil {
		return fmt.Errorf("parameter count is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateResourcePropertiesCountIsParameters(type_ *string, props interface{}, count *float64) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	if count == nil {
		return fmt.Errorf("parameter count is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Template) validateTemplateMatchesParameters(expected interface{}) error {
	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	return nil
}

func validateTemplate_FromJSONParameters(template *map[string]interface{}, templateParsingOptions *TemplateParsingOptions) error {
	if template == nil {
		return fmt.Errorf("parameter template is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(templateParsingOptions, func() string { return "parameter templateParsingOptions" }); err != nil {
		return err
	}

	return nil
}

func validateTemplate_FromStackParameters(stack awscdk.Stack, templateParsingOptions *TemplateParsingOptions) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(templateParsingOptions, func() string { return "parameter templateParsingOptions" }); err != nil {
		return err
	}

	return nil
}

func validateTemplate_FromStringParameters(template *string, templateParsingOptions *TemplateParsingOptions) error {
	if template == nil {
		return fmt.Errorf("parameter template is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(templateParsingOptions, func() string { return "parameter templateParsingOptions" }); err != nil {
		return err
	}

	return nil
}


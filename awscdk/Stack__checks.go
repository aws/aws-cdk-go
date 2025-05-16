//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_Stack) validateAddDependencyParameters(target Stack) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Stack) validateAddMetadataParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Stack) validateAddTransformParameters(transform *string) error {
	if transform == nil {
		return fmt.Errorf("parameter transform is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Stack) validateAllocateLogicalIdParameters(cfnElement CfnElement) error {
	if cfnElement == nil {
		return fmt.Errorf("parameter cfnElement is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Stack) validateExportStringListValueParameters(exportedValue interface{}, options *ExportValueOptions) error {
	if exportedValue == nil {
		return fmt.Errorf("parameter exportedValue is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_Stack) validateExportValueParameters(exportedValue interface{}, options *ExportValueOptions) error {
	if exportedValue == nil {
		return fmt.Errorf("parameter exportedValue is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_Stack) validateFormatArnParameters(components *ArnComponents) error {
	if components == nil {
		return fmt.Errorf("parameter components is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(components, func() string { return "parameter components" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_Stack) validateGetLogicalIdParameters(element CfnElement) error {
	if element == nil {
		return fmt.Errorf("parameter element is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Stack) validateRegionalFactParameters(factName *string) error {
	if factName == nil {
		return fmt.Errorf("parameter factName is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Stack) validateRenameLogicalIdParameters(oldId *string, newId *string) error {
	if oldId == nil {
		return fmt.Errorf("parameter oldId is required, but nil was provided")
	}

	if newId == nil {
		return fmt.Errorf("parameter newId is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Stack) validateReportMissingContextKeyParameters(report *cloudassemblyschema.MissingContext) error {
	if report == nil {
		return fmt.Errorf("parameter report is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(report, func() string { return "parameter report" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_Stack) validateResolveParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Stack) validateSplitArnParameters(arn *string, arnFormat ArnFormat) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	if arnFormat == "" {
		return fmt.Errorf("parameter arnFormat is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Stack) validateToJsonStringParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Stack) validateToYamlStringParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func validateStack_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateStack_IsStackParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateStack_OfParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Stack) validateSetTerminationProtectionParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewStackParameters(props *StackProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


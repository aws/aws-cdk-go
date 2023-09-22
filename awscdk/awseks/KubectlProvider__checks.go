//go:build !no_runtime_type_checking

package awseks

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
	"github.com/aws/constructs-go/constructs/v10"
)

func (k *jsiiProxy_KubectlProvider) validateAddDependencyParameters(target awscdk.Stack) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateAddMetadataParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateAddTransformParameters(transform *string) error {
	if transform == nil {
		return fmt.Errorf("parameter transform is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateAllocateLogicalIdParameters(cfnElement awscdk.CfnElement) error {
	if cfnElement == nil {
		return fmt.Errorf("parameter cfnElement is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateExportStringListValueParameters(exportedValue interface{}, options *awscdk.ExportValueOptions) error {
	if exportedValue == nil {
		return fmt.Errorf("parameter exportedValue is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateExportValueParameters(exportedValue interface{}, options *awscdk.ExportValueOptions) error {
	if exportedValue == nil {
		return fmt.Errorf("parameter exportedValue is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateFormatArnParameters(components *awscdk.ArnComponents) error {
	if components == nil {
		return fmt.Errorf("parameter components is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(components, func() string { return "parameter components" }); err != nil {
		return err
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateGetLogicalIdParameters(element awscdk.CfnElement) error {
	if element == nil {
		return fmt.Errorf("parameter element is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateRegionalFactParameters(factName *string) error {
	if factName == nil {
		return fmt.Errorf("parameter factName is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateRenameLogicalIdParameters(oldId *string, newId *string) error {
	if oldId == nil {
		return fmt.Errorf("parameter oldId is required, but nil was provided")
	}

	if newId == nil {
		return fmt.Errorf("parameter newId is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateReportMissingContextKeyParameters(report *cloudassemblyschema.MissingContext) error {
	if report == nil {
		return fmt.Errorf("parameter report is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(report, func() string { return "parameter report" }); err != nil {
		return err
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateResolveParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateSetParameterParameters(name *string, value *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateSplitArnParameters(arn *string, arnFormat awscdk.ArnFormat) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	if arnFormat == "" {
		return fmt.Errorf("parameter arnFormat is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateToJsonStringParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KubectlProvider) validateToYamlStringParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func validateKubectlProvider_FromKubectlProviderAttributesParameters(scope constructs.Construct, id *string, attrs *KubectlProviderAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validateKubectlProvider_GetOrCreateParameters(scope constructs.Construct, cluster ICluster) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if cluster == nil {
		return fmt.Errorf("parameter cluster is required, but nil was provided")
	}

	return nil
}

func validateKubectlProvider_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateKubectlProvider_IsNestedStackParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateKubectlProvider_IsStackParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateKubectlProvider_OfParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KubectlProvider) validateSetTerminationProtectionParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewKubectlProviderParameters(scope constructs.Construct, id *string, props *KubectlProviderProps) error {
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


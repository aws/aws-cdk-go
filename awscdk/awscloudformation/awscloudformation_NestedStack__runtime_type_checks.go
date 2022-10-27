//go:build !no_runtime_type_checking

package awscloudformation

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/cloudassemblyschema"
	"github.com/aws/aws-cdk-go/awscdk/cxapi"
	"github.com/aws/constructs-go/constructs/v3"
)

func (n *jsiiProxy_NestedStack) validateAddDependencyParameters(target awscdk.Stack) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateAddDockerImageAssetParameters(asset *awscdk.DockerImageAssetSource) error {
	if asset == nil {
		return fmt.Errorf("parameter asset is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(asset, func() string { return "parameter asset" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateAddFileAssetParameters(asset *awscdk.FileAssetSource) error {
	if asset == nil {
		return fmt.Errorf("parameter asset is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(asset, func() string { return "parameter asset" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateAddTransformParameters(transform *string) error {
	if transform == nil {
		return fmt.Errorf("parameter transform is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateAllocateLogicalIdParameters(cfnElement awscdk.CfnElement) error {
	if cfnElement == nil {
		return fmt.Errorf("parameter cfnElement is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateExportValueParameters(exportedValue interface{}, options *awscdk.ExportValueOptions) error {
	if exportedValue == nil {
		return fmt.Errorf("parameter exportedValue is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateFormatArnParameters(components *awscdk.ArnComponents) error {
	if components == nil {
		return fmt.Errorf("parameter components is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(components, func() string { return "parameter components" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateGetLogicalIdParameters(element awscdk.CfnElement) error {
	if element == nil {
		return fmt.Errorf("parameter element is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateParseArnParameters(arn *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validatePrepareCrossReferenceParameters(_sourceStack awscdk.Stack, reference awscdk.Reference) error {
	if _sourceStack == nil {
		return fmt.Errorf("parameter _sourceStack is required, but nil was provided")
	}

	if reference == nil {
		return fmt.Errorf("parameter reference is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateRegionalFactParameters(factName *string) error {
	if factName == nil {
		return fmt.Errorf("parameter factName is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateRenameLogicalIdParameters(oldId *string, newId *string) error {
	if oldId == nil {
		return fmt.Errorf("parameter oldId is required, but nil was provided")
	}

	if newId == nil {
		return fmt.Errorf("parameter newId is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateReportMissingContextParameters(report *cxapi.MissingContext) error {
	if report == nil {
		return fmt.Errorf("parameter report is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(report, func() string { return "parameter report" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateReportMissingContextKeyParameters(report *cloudassemblyschema.MissingContext) error {
	if report == nil {
		return fmt.Errorf("parameter report is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(report, func() string { return "parameter report" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateResolveParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateSetParameterParameters(name *string, value *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateSplitArnParameters(arn *string, arnFormat awscdk.ArnFormat) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	if arnFormat == "" {
		return fmt.Errorf("parameter arnFormat is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStack) validateToJsonStringParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func validateNestedStack_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNestedStack_IsNestedStackParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNestedStack_IsStackParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNestedStack_OfParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewNestedStackParameters(scope awscdk.Construct, id *string, props *NestedStackProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


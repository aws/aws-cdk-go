//go:build !no_runtime_type_checking

package awsbedrockalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (g *jsiiProxy_Guardrail) validateAddContentFilterParameters(filter *ContentFilter) error {
	if filter == nil {
		return fmt.Errorf("parameter filter is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(filter, func() string { return "parameter filter" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateAddContextualGroundingFilterParameters(filter *ContextualGroundingFilter) error {
	if filter == nil {
		return fmt.Errorf("parameter filter is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(filter, func() string { return "parameter filter" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateAddDeniedTopicFilterParameters(filter Topic) error {
	if filter == nil {
		return fmt.Errorf("parameter filter is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateAddManagedWordListFilterParameters(filter *ManagedWordFilter) error {
	if filter == nil {
		return fmt.Errorf("parameter filter is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(filter, func() string { return "parameter filter" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateAddPIIFilterParameters(filter *PIIFilter) error {
	if filter == nil {
		return fmt.Errorf("parameter filter is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(filter, func() string { return "parameter filter" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateAddRegexFilterParameters(filter *RegexFilter) error {
	if filter == nil {
		return fmt.Errorf("parameter filter is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(filter, func() string { return "parameter filter" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateAddWordFilterParameters(filter *WordFilter) error {
	if filter == nil {
		return fmt.Errorf("parameter filter is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(filter, func() string { return "parameter filter" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateAddWordFilterFromFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateGrantParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateGrantApplyParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateMetricInvocationClientErrorsParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateMetricInvocationLatencyParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateMetricInvocationServerErrorsParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateMetricInvocationsIntervenedParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateMetricInvocationThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateMetricTextUnitCountParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateUpdateVersionParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

func validateGuardrail_FromCfnGuardrailParameters(cfnGuardrail awsbedrock.CfnGuardrail) error {
	if cfnGuardrail == nil {
		return fmt.Errorf("parameter cfnGuardrail is required, but nil was provided")
	}

	return nil
}

func validateGuardrail_FromGuardrailAttributesParameters(scope constructs.Construct, id *string, attrs *GuardrailAttributes) error {
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

func validateGuardrail_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateGuardrail_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateGuardrail_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateGuardrail_MetricAllParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateGuardrail_MetricAllInvocationLatencyParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateGuardrail_MetricAllInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateGuardrail_MetricAllInvocationsIntervenedParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateGuardrail_MetricAllTextUnitCountParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNewGuardrailParameters(scope constructs.Construct, id *string, props *GuardrailProps) error {
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


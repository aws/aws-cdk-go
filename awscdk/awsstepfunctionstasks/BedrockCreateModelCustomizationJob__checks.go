//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	if branch == nil {
		return fmt.Errorf("parameter branch is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	if handler == nil {
		return fmt.Errorf("parameter handler is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State, options *awsstepfunctions.ChoiceTransitionOptions) error {
	if condition == nil {
		return fmt.Errorf("parameter condition is required, but nil was provided")
	}

	if next == nil {
		return fmt.Errorf("parameter next is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateAddItemProcessorParameters(processor awsstepfunctions.StateGraph, config *awsstepfunctions.ProcessorConfig) error {
	if processor == nil {
		return fmt.Errorf("parameter processor is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	if iteration == nil {
		return fmt.Errorf("parameter iteration is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateAddPrefixParameters(x *string) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	if graph == nil {
		return fmt.Errorf("parameter graph is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	if def == nil {
		return fmt.Errorf("parameter def is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateMakeNextParameters(next awsstepfunctions.State) error {
	if next == nil {
		return fmt.Errorf("parameter next is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateNextParameters(next awsstepfunctions.IChainable) error {
	if next == nil {
		return fmt.Errorf("parameter next is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	if graph == nil {
		return fmt.Errorf("parameter graph is required, but nil was provided")
	}

	return nil
}

func validateBedrockCreateModelCustomizationJob_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	if states == nil {
		return fmt.Errorf("parameter states is required, but nil was provided")
	}

	return nil
}

func validateBedrockCreateModelCustomizationJob_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateBedrockCreateModelCustomizationJob_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateBedrockCreateModelCustomizationJob_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateBedrockCreateModelCustomizationJob_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	if root == nil {
		return fmt.Errorf("parameter root is required, but nil was provided")
	}

	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) validateSetProcessorConfigParameters(val *awsstepfunctions.ProcessorConfig) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func validateNewBedrockCreateModelCustomizationJobParameters(scope constructs.Construct, id *string, props *BedrockCreateModelCustomizationJobProps) error {
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


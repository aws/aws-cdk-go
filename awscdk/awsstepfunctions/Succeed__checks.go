//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_Succeed) validateAddBranchParameters(branch StateGraph) error {
	if branch == nil {
		return fmt.Errorf("parameter branch is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Succeed) validateAddChoiceParameters(condition Condition, next State, options *ChoiceTransitionOptions) error {
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

func (s *jsiiProxy_Succeed) validateAddItemProcessorParameters(processor StateGraph, config *ProcessorConfig) error {
	if processor == nil {
		return fmt.Errorf("parameter processor is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_Succeed) validateAddIteratorParameters(iteration StateGraph) error {
	if iteration == nil {
		return fmt.Errorf("parameter iteration is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Succeed) validateAddPrefixParameters(x *string) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Succeed) validateBindToGraphParameters(graph StateGraph) error {
	if graph == nil {
		return fmt.Errorf("parameter graph is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Succeed) validateMakeDefaultParameters(def State) error {
	if def == nil {
		return fmt.Errorf("parameter def is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Succeed) validateMakeNextParameters(next State) error {
	if next == nil {
		return fmt.Errorf("parameter next is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Succeed) validateWhenBoundToGraphParameters(graph StateGraph) error {
	if graph == nil {
		return fmt.Errorf("parameter graph is required, but nil was provided")
	}

	return nil
}

func validateSucceed_FilterNextablesParameters(states *[]State) error {
	if states == nil {
		return fmt.Errorf("parameter states is required, but nil was provided")
	}

	return nil
}

func validateSucceed_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateSucceed_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateSucceed_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateSucceed_JsonataParameters(scope constructs.Construct, id *string, props *SucceedJsonataProps) error {
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

func validateSucceed_JsonPathParameters(scope constructs.Construct, id *string, props *SucceedJsonPathProps) error {
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

func validateSucceed_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	if root == nil {
		return fmt.Errorf("parameter root is required, but nil was provided")
	}

	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Succeed) validateSetProcessorConfigParameters(val *ProcessorConfig) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func validateNewSucceedParameters(scope constructs.Construct, id *string, props *SucceedProps) error {
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


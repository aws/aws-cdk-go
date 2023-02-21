//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_Choice) validateAddBranchParameters(branch StateGraph) error {
	if branch == nil {
		return fmt.Errorf("parameter branch is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Choice) validateAddChoiceParameters(condition Condition, next State) error {
	if condition == nil {
		return fmt.Errorf("parameter condition is required, but nil was provided")
	}

	if next == nil {
		return fmt.Errorf("parameter next is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Choice) validateAddIteratorParameters(iteration StateGraph) error {
	if iteration == nil {
		return fmt.Errorf("parameter iteration is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Choice) validateAddPrefixParameters(x *string) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Choice) validateAfterwardsParameters(options *AfterwardsOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_Choice) validateBindToGraphParameters(graph StateGraph) error {
	if graph == nil {
		return fmt.Errorf("parameter graph is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Choice) validateMakeDefaultParameters(def State) error {
	if def == nil {
		return fmt.Errorf("parameter def is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Choice) validateMakeNextParameters(next State) error {
	if next == nil {
		return fmt.Errorf("parameter next is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Choice) validateOtherwiseParameters(def IChainable) error {
	if def == nil {
		return fmt.Errorf("parameter def is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Choice) validateWhenParameters(condition Condition, next IChainable) error {
	if condition == nil {
		return fmt.Errorf("parameter condition is required, but nil was provided")
	}

	if next == nil {
		return fmt.Errorf("parameter next is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Choice) validateWhenBoundToGraphParameters(graph StateGraph) error {
	if graph == nil {
		return fmt.Errorf("parameter graph is required, but nil was provided")
	}

	return nil
}

func validateChoice_FilterNextablesParameters(states *[]State) error {
	if states == nil {
		return fmt.Errorf("parameter states is required, but nil was provided")
	}

	return nil
}

func validateChoice_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateChoice_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateChoice_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateChoice_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	if root == nil {
		return fmt.Errorf("parameter root is required, but nil was provided")
	}

	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func validateNewChoiceParameters(scope constructs.Construct, id *string, props *ChoiceProps) error {
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


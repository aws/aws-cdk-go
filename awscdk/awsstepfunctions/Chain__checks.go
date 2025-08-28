//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (c *jsiiProxy_Chain) validateNextParameters(next IChainable) error {
	if next == nil {
		return fmt.Errorf("parameter next is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Chain) validateToSingleStateParameters(id *string, props *ParallelProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateChain_CustomParameters(startState State, endStates *[]INextable, lastAdded IChainable) error {
	if startState == nil {
		return fmt.Errorf("parameter startState is required, but nil was provided")
	}

	if endStates == nil {
		return fmt.Errorf("parameter endStates is required, but nil was provided")
	}

	if lastAdded == nil {
		return fmt.Errorf("parameter lastAdded is required, but nil was provided")
	}

	return nil
}

func validateChain_SequenceParameters(start IChainable, next IChainable) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if next == nil {
		return fmt.Errorf("parameter next is required, but nil was provided")
	}

	return nil
}

func validateChain_StartParameters(state IChainable) error {
	if state == nil {
		return fmt.Errorf("parameter state is required, but nil was provided")
	}

	return nil
}


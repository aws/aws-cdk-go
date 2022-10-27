//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (c *jsiiProxy_Connections) validateAllowDefaultPortFromParameters(other IConnectable) error {
	if other == nil {
		return fmt.Errorf("parameter other is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connections) validateAllowDefaultPortToParameters(other IConnectable) error {
	if other == nil {
		return fmt.Errorf("parameter other is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connections) validateAllowFromParameters(other IConnectable, portRange Port) error {
	if other == nil {
		return fmt.Errorf("parameter other is required, but nil was provided")
	}

	if portRange == nil {
		return fmt.Errorf("parameter portRange is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connections) validateAllowFromAnyIpv4Parameters(portRange Port) error {
	if portRange == nil {
		return fmt.Errorf("parameter portRange is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connections) validateAllowInternallyParameters(portRange Port) error {
	if portRange == nil {
		return fmt.Errorf("parameter portRange is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connections) validateAllowToParameters(other IConnectable, portRange Port) error {
	if other == nil {
		return fmt.Errorf("parameter other is required, but nil was provided")
	}

	if portRange == nil {
		return fmt.Errorf("parameter portRange is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connections) validateAllowToAnyIpv4Parameters(portRange Port) error {
	if portRange == nil {
		return fmt.Errorf("parameter portRange is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connections) validateAllowToDefaultPortParameters(other IConnectable) error {
	if other == nil {
		return fmt.Errorf("parameter other is required, but nil was provided")
	}

	return nil
}

func validateNewConnectionsParameters(props *ConnectionsProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validatePort_TcpParameters(port *float64) error {
	if port == nil {
		return fmt.Errorf("parameter port is required, but nil was provided")
	}

	return nil
}

func validatePort_TcpRangeParameters(startPort *float64, endPort *float64) error {
	if startPort == nil {
		return fmt.Errorf("parameter startPort is required, but nil was provided")
	}

	if endPort == nil {
		return fmt.Errorf("parameter endPort is required, but nil was provided")
	}

	return nil
}

func validatePort_UdpParameters(port *float64) error {
	if port == nil {
		return fmt.Errorf("parameter port is required, but nil was provided")
	}

	return nil
}

func validatePort_UdpRangeParameters(startPort *float64, endPort *float64) error {
	if startPort == nil {
		return fmt.Errorf("parameter startPort is required, but nil was provided")
	}

	if endPort == nil {
		return fmt.Errorf("parameter endPort is required, but nil was provided")
	}

	return nil
}

func validateNewPortParameters(props *PortProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


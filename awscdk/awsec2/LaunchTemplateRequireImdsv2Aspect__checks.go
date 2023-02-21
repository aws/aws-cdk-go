//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_LaunchTemplateRequireImdsv2Aspect) validateVisitParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LaunchTemplateRequireImdsv2Aspect) validateWarnParameters(node constructs.IConstruct, message *string) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func validateNewLaunchTemplateRequireImdsv2AspectParameters(props *LaunchTemplateRequireImdsv2AspectProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


//go:build !no_runtime_type_checking

package awsiotfleetwise

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CfnStateTemplatePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStateTemplatePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateCfnStateTemplatePropsMixin_IsMixinParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewCfnStateTemplatePropsMixinParameters(props *CfnStateTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"
)

func (d *jsiiProxy_DefaultTokenResolver) validateResolveListParameters(xs *[]*string, context IResolveContext) error {
	if xs == nil {
		return fmt.Errorf("parameter xs is required, but nil was provided")
	}

	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DefaultTokenResolver) validateResolveStringParameters(fragments TokenizedStringFragments, context IResolveContext) error {
	if fragments == nil {
		return fmt.Errorf("parameter fragments is required, but nil was provided")
	}

	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DefaultTokenResolver) validateResolveTokenParameters(t IResolvable, context IResolveContext, postProcessor IPostProcessor) error {
	if t == nil {
		return fmt.Errorf("parameter t is required, but nil was provided")
	}

	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	if postProcessor == nil {
		return fmt.Errorf("parameter postProcessor is required, but nil was provided")
	}

	return nil
}

func validateNewDefaultTokenResolverParameters(concat IFragmentConcatenator) error {
	if concat == nil {
		return fmt.Errorf("parameter concat is required, but nil was provided")
	}

	return nil
}


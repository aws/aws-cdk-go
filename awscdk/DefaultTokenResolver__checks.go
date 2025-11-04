//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func (d *jsiiProxy_DefaultTokenResolver) validateResolveListParameters(l *[]*string, context IResolveContext) error {
	if l == nil {
		return fmt.Errorf("parameter l is required, but nil was provided")
	}

	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DefaultTokenResolver) validateResolveStringParameters(s TokenizedStringFragments, context IResolveContext) error {
	if s == nil {
		return fmt.Errorf("parameter s is required, but nil was provided")
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


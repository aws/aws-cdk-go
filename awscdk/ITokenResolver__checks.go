//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"
)

func (i *jsiiProxy_ITokenResolver) validateResolveListParameters(l *[]*string, context IResolveContext) error {
	if l == nil {
		return fmt.Errorf("parameter l is required, but nil was provided")
	}

	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ITokenResolver) validateResolveStringParameters(s TokenizedStringFragments, context IResolveContext) error {
	if s == nil {
		return fmt.Errorf("parameter s is required, but nil was provided")
	}

	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ITokenResolver) validateResolveTokenParameters(t IResolvable, context IResolveContext, postProcessor IPostProcessor) error {
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


//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DefaultTokenResolver) validateResolveListParameters(xs *[]*string, context IResolveContext) error {
	return nil
}

func (d *jsiiProxy_DefaultTokenResolver) validateResolveStringParameters(fragments TokenizedStringFragments, context IResolveContext) error {
	return nil
}

func (d *jsiiProxy_DefaultTokenResolver) validateResolveTokenParameters(t IResolvable, context IResolveContext, postProcessor IPostProcessor) error {
	return nil
}

func validateNewDefaultTokenResolverParameters(concat IFragmentConcatenator) error {
	return nil
}


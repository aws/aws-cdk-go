//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TokenizedStringFragments) validateAddIntrinsicParameters(value interface{}) error {
	return nil
}

func (t *jsiiProxy_TokenizedStringFragments) validateAddLiteralParameters(lit interface{}) error {
	return nil
}

func (t *jsiiProxy_TokenizedStringFragments) validateAddTokenParameters(token IResolvable) error {
	return nil
}

func (t *jsiiProxy_TokenizedStringFragments) validateJoinParameters(concat IFragmentConcatenator) error {
	return nil
}

func (t *jsiiProxy_TokenizedStringFragments) validateMapTokensParameters(mapper ITokenMapper) error {
	return nil
}


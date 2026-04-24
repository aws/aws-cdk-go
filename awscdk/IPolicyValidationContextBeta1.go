package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Context available to the validation plugin.
// Deprecated: Use `IPolicyValidationContext` instead.
type IPolicyValidationContextBeta1 interface {
	// The absolute path of all templates to be processed.
	// Deprecated: Use `IPolicyValidationContext` instead.
	TemplatePaths() *[]*string
}

// The jsii proxy for IPolicyValidationContextBeta1
type jsiiProxy_IPolicyValidationContextBeta1 struct {
	_ byte // padding
}

func (j *jsiiProxy_IPolicyValidationContextBeta1) TemplatePaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"templatePaths",
		&returns,
	)
	return returns
}


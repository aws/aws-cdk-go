package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Context available to the validation plugin.
type IPolicyValidationContext interface {
	// The root construct of the app being validated.
	//
	// Plugins may walk this tree for typed L1 property access and token
	// resolution via `Stack.of(node).resolve()`. The tree is finalized and
	// should be treated as read-only; mutations have no effect on synthesized
	// output.
	AppConstruct() constructs.IConstruct
	// The absolute path of all templates to be processed.
	TemplatePaths() *[]*string
}

// The jsii proxy for IPolicyValidationContext
type jsiiProxy_IPolicyValidationContext struct {
	_ byte // padding
}

func (j *jsiiProxy_IPolicyValidationContext) AppConstruct() constructs.IConstruct {
	var returns constructs.IConstruct
	_jsii_.Get(
		j,
		"appConstruct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyValidationContext) TemplatePaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"templatePaths",
		&returns,
	)
	return returns
}


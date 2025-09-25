package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswaf/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SqlInjectionMatchSet.
// Experimental.
type ISqlInjectionMatchSetRef interface {
	constructs.IConstruct
	// A reference to a SqlInjectionMatchSet resource.
	// Experimental.
	SqlInjectionMatchSetRef() *SqlInjectionMatchSetReference
}

// The jsii proxy for ISqlInjectionMatchSetRef
type jsiiProxy_ISqlInjectionMatchSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISqlInjectionMatchSetRef) SqlInjectionMatchSetRef() *SqlInjectionMatchSetReference {
	var returns *SqlInjectionMatchSetReference
	_jsii_.Get(
		j,
		"sqlInjectionMatchSetRef",
		&returns,
	)
	return returns
}


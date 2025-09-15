package awswafregional

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswafregional/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RegexPatternSet.
// Experimental.
type IRegexPatternSetRef interface {
	constructs.IConstruct
	// A reference to a RegexPatternSet resource.
	// Experimental.
	RegexPatternSetRef() *RegexPatternSetReference
}

// The jsii proxy for IRegexPatternSetRef
type jsiiProxy_IRegexPatternSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRegexPatternSetRef) RegexPatternSetRef() *RegexPatternSetReference {
	var returns *RegexPatternSetReference
	_jsii_.Get(
		j,
		"regexPatternSetRef",
		&returns,
	)
	return returns
}


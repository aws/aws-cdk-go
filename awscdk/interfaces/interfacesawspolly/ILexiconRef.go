package interfacesawspolly

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspolly/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Lexicon.
// Experimental.
type ILexiconRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Lexicon resource.
	// Experimental.
	LexiconRef() *LexiconReference
}

// The jsii proxy for ILexiconRef
type jsiiProxy_ILexiconRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILexiconRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ILexiconRef) LexiconRef() *LexiconReference {
	var returns *LexiconReference
	_jsii_.Get(
		j,
		"lexiconRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILexiconRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILexiconRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package interfacesawskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskendra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Thesaurus.
// Experimental.
type IThesaurusRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Thesaurus resource.
	// Experimental.
	ThesaurusRef() *ThesaurusReference
}

// The jsii proxy for IThesaurusRef
type jsiiProxy_IThesaurusRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IThesaurusRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IThesaurusRef) ThesaurusRef() *ThesaurusReference {
	var returns *ThesaurusReference
	_jsii_.Get(
		j,
		"thesaurusRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IThesaurusRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IThesaurusRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


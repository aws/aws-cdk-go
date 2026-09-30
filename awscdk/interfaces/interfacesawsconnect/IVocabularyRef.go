package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Vocabulary.
// Experimental.
type IVocabularyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Vocabulary resource.
	// Experimental.
	VocabularyRef() *VocabularyReference
}

// The jsii proxy for IVocabularyRef
type jsiiProxy_IVocabularyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IVocabularyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IVocabularyRef) VocabularyRef() *VocabularyReference {
	var returns *VocabularyReference
	_jsii_.Get(
		j,
		"vocabularyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVocabularyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVocabularyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


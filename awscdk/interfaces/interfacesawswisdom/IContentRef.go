package interfacesawswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Content.
// Experimental.
type IContentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Content resource.
	// Experimental.
	ContentRef() *ContentReference
}

// The jsii proxy for IContentRef
type jsiiProxy_IContentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IContentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IContentRef) ContentRef() *ContentReference {
	var returns *ContentReference
	_jsii_.Get(
		j,
		"contentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


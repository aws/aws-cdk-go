package interfacesawscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Publisher.
// Experimental.
type IPublisherRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Publisher resource.
	// Experimental.
	PublisherRef() *PublisherReference
}

// The jsii proxy for IPublisherRef
type jsiiProxy_IPublisherRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPublisherRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPublisherRef) PublisherRef() *PublisherReference {
	var returns *PublisherReference
	_jsii_.Get(
		j,
		"publisherRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublisherRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublisherRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


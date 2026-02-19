package interfacesawscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PublicTypeVersion.
// Experimental.
type IPublicTypeVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PublicTypeVersion resource.
	// Experimental.
	PublicTypeVersionRef() *PublicTypeVersionReference
}

// The jsii proxy for IPublicTypeVersionRef
type jsiiProxy_IPublicTypeVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPublicTypeVersionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPublicTypeVersionRef) PublicTypeVersionRef() *PublicTypeVersionReference {
	var returns *PublicTypeVersionReference
	_jsii_.Get(
		j,
		"publicTypeVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicTypeVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicTypeVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


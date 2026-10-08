package interfacesawscomprehend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscomprehend/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EntityRecognizerEndpoint.
// Experimental.
type IEntityRecognizerEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EntityRecognizerEndpoint resource.
	// Experimental.
	EntityRecognizerEndpointRef() *EntityRecognizerEndpointReference
}

// The jsii proxy for IEntityRecognizerEndpointRef
type jsiiProxy_IEntityRecognizerEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEntityRecognizerEndpointRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEntityRecognizerEndpointRef) EntityRecognizerEndpointRef() *EntityRecognizerEndpointReference {
	var returns *EntityRecognizerEndpointReference
	_jsii_.Get(
		j,
		"entityRecognizerEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEntityRecognizerEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEntityRecognizerEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


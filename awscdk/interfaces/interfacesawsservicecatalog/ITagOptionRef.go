package interfacesawsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TagOption.
// Experimental.
type ITagOptionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TagOption resource.
	// Experimental.
	TagOptionRef() *TagOptionReference
}

// The jsii proxy for ITagOptionRef
type jsiiProxy_ITagOptionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITagOptionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITagOptionRef) TagOptionRef() *TagOptionReference {
	var returns *TagOptionReference
	_jsii_.Get(
		j,
		"tagOptionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITagOptionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITagOptionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


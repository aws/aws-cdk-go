package interfacesawscases

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscases/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Layout.
// Experimental.
type ILayoutRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Layout resource.
	// Experimental.
	LayoutRef() *LayoutReference
}

// The jsii proxy for ILayoutRef
type jsiiProxy_ILayoutRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILayoutRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ILayoutRef) LayoutRef() *LayoutReference {
	var returns *LayoutReference
	_jsii_.Get(
		j,
		"layoutRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILayoutRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILayoutRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


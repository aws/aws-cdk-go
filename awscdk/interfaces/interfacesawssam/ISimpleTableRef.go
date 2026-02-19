package interfacesawssam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SimpleTable.
// Experimental.
type ISimpleTableRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SimpleTable resource.
	// Experimental.
	SimpleTableRef() *SimpleTableReference
}

// The jsii proxy for ISimpleTableRef
type jsiiProxy_ISimpleTableRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISimpleTableRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISimpleTableRef) SimpleTableRef() *SimpleTableReference {
	var returns *SimpleTableReference
	_jsii_.Get(
		j,
		"simpleTableRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISimpleTableRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISimpleTableRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


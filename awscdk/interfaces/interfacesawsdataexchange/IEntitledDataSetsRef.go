package interfacesawsdataexchange

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdataexchange/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EntitledDataSets.
// Experimental.
type IEntitledDataSetsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EntitledDataSets resource.
	// Experimental.
	EntitledDataSetsRef() *EntitledDataSetsReference
}

// The jsii proxy for IEntitledDataSetsRef
type jsiiProxy_IEntitledDataSetsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEntitledDataSetsRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEntitledDataSetsRef) EntitledDataSetsRef() *EntitledDataSetsReference {
	var returns *EntitledDataSetsReference
	_jsii_.Get(
		j,
		"entitledDataSetsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEntitledDataSetsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEntitledDataSetsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


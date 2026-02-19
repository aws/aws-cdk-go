package interfacesawsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BatchScramSecret.
// Experimental.
type IBatchScramSecretRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a BatchScramSecret resource.
	// Experimental.
	BatchScramSecretRef() *BatchScramSecretReference
}

// The jsii proxy for IBatchScramSecretRef
type jsiiProxy_IBatchScramSecretRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IBatchScramSecretRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IBatchScramSecretRef) BatchScramSecretRef() *BatchScramSecretReference {
	var returns *BatchScramSecretReference
	_jsii_.Get(
		j,
		"batchScramSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBatchScramSecretRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBatchScramSecretRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


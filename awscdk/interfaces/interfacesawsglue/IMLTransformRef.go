package interfacesawsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MLTransform.
// Experimental.
type IMLTransformRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MLTransform resource.
	// Experimental.
	MlTransformRef() *MLTransformReference
}

// The jsii proxy for IMLTransformRef
type jsiiProxy_IMLTransformRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IMLTransformRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IMLTransformRef) MlTransformRef() *MLTransformReference {
	var returns *MLTransformReference
	_jsii_.Get(
		j,
		"mlTransformRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMLTransformRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMLTransformRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


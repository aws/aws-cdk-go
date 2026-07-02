package interfacesawslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MicrovmImage.
// Experimental.
type IMicrovmImageRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MicrovmImage resource.
	// Experimental.
	MicrovmImageRef() *MicrovmImageReference
}

// The jsii proxy for IMicrovmImageRef
type jsiiProxy_IMicrovmImageRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IMicrovmImageRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IMicrovmImageRef) MicrovmImageRef() *MicrovmImageReference {
	var returns *MicrovmImageReference
	_jsii_.Get(
		j,
		"microvmImageRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMicrovmImageRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMicrovmImageRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


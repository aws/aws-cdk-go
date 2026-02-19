package interfacesawsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FuotaTask.
// Experimental.
type IFuotaTaskRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FuotaTask resource.
	// Experimental.
	FuotaTaskRef() *FuotaTaskReference
}

// The jsii proxy for IFuotaTaskRef
type jsiiProxy_IFuotaTaskRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IFuotaTaskRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IFuotaTaskRef) FuotaTaskRef() *FuotaTaskReference {
	var returns *FuotaTaskReference
	_jsii_.Get(
		j,
		"fuotaTaskRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFuotaTaskRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFuotaTaskRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


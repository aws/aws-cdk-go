package interfacesawsthinclient

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsthinclient/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SoftwareSet.
// Experimental.
type ISoftwareSetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SoftwareSet resource.
	// Experimental.
	SoftwareSetRef() *SoftwareSetReference
}

// The jsii proxy for ISoftwareSetRef
type jsiiProxy_ISoftwareSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISoftwareSetRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISoftwareSetRef) SoftwareSetRef() *SoftwareSetReference {
	var returns *SoftwareSetReference
	_jsii_.Get(
		j,
		"softwareSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISoftwareSetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISoftwareSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


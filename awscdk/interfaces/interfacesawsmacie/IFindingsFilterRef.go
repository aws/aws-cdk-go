package interfacesawsmacie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmacie/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FindingsFilter.
// Experimental.
type IFindingsFilterRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FindingsFilter resource.
	// Experimental.
	FindingsFilterRef() *FindingsFilterReference
}

// The jsii proxy for IFindingsFilterRef
type jsiiProxy_IFindingsFilterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IFindingsFilterRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IFindingsFilterRef) FindingsFilterRef() *FindingsFilterReference {
	var returns *FindingsFilterReference
	_jsii_.Get(
		j,
		"findingsFilterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFindingsFilterRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFindingsFilterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


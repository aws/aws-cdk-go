package interfacesawspi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspi/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PerfReports.
// Experimental.
type IPerfReportsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PerfReports resource.
	// Experimental.
	PerfReportsRef() *PerfReportsReference
}

// The jsii proxy for IPerfReportsRef
type jsiiProxy_IPerfReportsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPerfReportsRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPerfReportsRef) PerfReportsRef() *PerfReportsReference {
	var returns *PerfReportsReference
	_jsii_.Get(
		j,
		"perfReportsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPerfReportsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPerfReportsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


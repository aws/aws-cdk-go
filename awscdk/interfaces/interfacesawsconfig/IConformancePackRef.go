package interfacesawsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConformancePack.
// Experimental.
type IConformancePackRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConformancePack resource.
	// Experimental.
	ConformancePackRef() *ConformancePackReference
}

// The jsii proxy for IConformancePackRef
type jsiiProxy_IConformancePackRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IConformancePackRef) ConformancePackRef() *ConformancePackReference {
	var returns *ConformancePackReference
	_jsii_.Get(
		j,
		"conformancePackRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConformancePackRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConformancePackRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


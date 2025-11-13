package interfacesawssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FindingAggregator.
// Experimental.
type IFindingAggregatorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FindingAggregator resource.
	// Experimental.
	FindingAggregatorRef() *FindingAggregatorReference
}

// The jsii proxy for IFindingAggregatorRef
type jsiiProxy_IFindingAggregatorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IFindingAggregatorRef) FindingAggregatorRef() *FindingAggregatorReference {
	var returns *FindingAggregatorReference
	_jsii_.Get(
		j,
		"findingAggregatorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFindingAggregatorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFindingAggregatorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


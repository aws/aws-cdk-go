package interfacesawssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AggregatorV2.
// Experimental.
type IAggregatorV2Ref interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AggregatorV2 resource.
	// Experimental.
	AggregatorV2Ref() *AggregatorV2Reference
}

// The jsii proxy for IAggregatorV2Ref
type jsiiProxy_IAggregatorV2Ref struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAggregatorV2Ref) AggregatorV2Ref() *AggregatorV2Reference {
	var returns *AggregatorV2Reference
	_jsii_.Get(
		j,
		"aggregatorV2Ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAggregatorV2Ref) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAggregatorV2Ref) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package interfacesawscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Distribution.
// Experimental.
type IDistributionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Distribution resource.
	// Experimental.
	DistributionRef() *DistributionReference
}

// The jsii proxy for IDistributionRef
type jsiiProxy_IDistributionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDistributionRef) DistributionRef() *DistributionReference {
	var returns *DistributionReference
	_jsii_.Get(
		j,
		"distributionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDistributionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDistributionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


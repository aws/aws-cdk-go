package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrafficDistributionGroup.
// Experimental.
type ITrafficDistributionGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TrafficDistributionGroup resource.
	// Experimental.
	TrafficDistributionGroupRef() *TrafficDistributionGroupReference
}

// The jsii proxy for ITrafficDistributionGroupRef
type jsiiProxy_ITrafficDistributionGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ITrafficDistributionGroupRef) TrafficDistributionGroupRef() *TrafficDistributionGroupReference {
	var returns *TrafficDistributionGroupReference
	_jsii_.Get(
		j,
		"trafficDistributionGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrafficDistributionGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrafficDistributionGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


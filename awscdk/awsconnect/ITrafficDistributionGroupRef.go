package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrafficDistributionGroup.
// Experimental.
type ITrafficDistributionGroupRef interface {
	constructs.IConstruct
	// A reference to a TrafficDistributionGroup resource.
	// Experimental.
	TrafficDistributionGroupRef() *TrafficDistributionGroupReference
}

// The jsii proxy for ITrafficDistributionGroupRef
type jsiiProxy_ITrafficDistributionGroupRef struct {
	internal.Type__constructsIConstruct
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


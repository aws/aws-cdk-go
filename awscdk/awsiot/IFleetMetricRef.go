package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FleetMetric.
// Experimental.
type IFleetMetricRef interface {
	constructs.IConstruct
	// A reference to a FleetMetric resource.
	// Experimental.
	FleetMetricRef() *FleetMetricReference
}

// The jsii proxy for IFleetMetricRef
type jsiiProxy_IFleetMetricRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFleetMetricRef) FleetMetricRef() *FleetMetricReference {
	var returns *FleetMetricReference
	_jsii_.Get(
		j,
		"fleetMetricRef",
		&returns,
	)
	return returns
}


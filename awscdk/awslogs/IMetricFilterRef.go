package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MetricFilter.
// Experimental.
type IMetricFilterRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMetricFilterRef
type jsiiProxy_IMetricFilterRef struct {
	internal.Type__constructsIConstruct
}


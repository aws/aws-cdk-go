package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MetricStream.
// Experimental.
type IMetricStreamRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMetricStreamRef
type jsiiProxy_IMetricStreamRef struct {
	internal.Type__constructsIConstruct
}


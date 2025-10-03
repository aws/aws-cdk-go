package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomMetric.
// Experimental.
type ICustomMetricRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICustomMetricRef
type jsiiProxy_ICustomMetricRef struct {
	internal.Type__constructsIConstruct
}


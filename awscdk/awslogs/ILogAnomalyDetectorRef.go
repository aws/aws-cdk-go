package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogAnomalyDetector.
// Experimental.
type ILogAnomalyDetectorRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILogAnomalyDetectorRef
type jsiiProxy_ILogAnomalyDetectorRef struct {
	internal.Type__constructsIConstruct
}


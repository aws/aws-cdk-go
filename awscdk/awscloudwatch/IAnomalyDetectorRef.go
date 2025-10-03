package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnomalyDetector.
// Experimental.
type IAnomalyDetectorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAnomalyDetectorRef
type jsiiProxy_IAnomalyDetectorRef struct {
	internal.Type__constructsIConstruct
}


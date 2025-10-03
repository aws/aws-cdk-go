package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MlflowTrackingServer.
// Experimental.
type IMlflowTrackingServerRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMlflowTrackingServerRef
type jsiiProxy_IMlflowTrackingServerRef struct {
	internal.Type__constructsIConstruct
}


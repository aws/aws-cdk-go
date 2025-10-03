package awsm2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsm2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Deployment.
// Experimental.
type IDeploymentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDeploymentRef
type jsiiProxy_IDeploymentRef struct {
	internal.Type__constructsIConstruct
}


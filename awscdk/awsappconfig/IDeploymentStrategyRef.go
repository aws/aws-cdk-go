package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeploymentStrategy.
// Experimental.
type IDeploymentStrategyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDeploymentStrategyRef
type jsiiProxy_IDeploymentStrategyRef struct {
	internal.Type__constructsIConstruct
}


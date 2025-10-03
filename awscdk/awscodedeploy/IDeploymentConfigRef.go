package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeploymentConfig.
// Experimental.
type IDeploymentConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDeploymentConfigRef
type jsiiProxy_IDeploymentConfigRef struct {
	internal.Type__constructsIConstruct
}


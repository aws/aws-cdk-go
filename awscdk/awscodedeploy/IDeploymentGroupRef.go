package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeploymentGroup.
// Experimental.
type IDeploymentGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDeploymentGroupRef
type jsiiProxy_IDeploymentGroupRef struct {
	internal.Type__constructsIConstruct
}


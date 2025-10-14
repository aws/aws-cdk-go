package awsappconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeploymentStrategy.
// Experimental.
type IDeploymentStrategyRef interface {
	constructs.IConstruct
	// A reference to a DeploymentStrategy resource.
	// Experimental.
	DeploymentStrategyRef() *DeploymentStrategyReference
}

// The jsii proxy for IDeploymentStrategyRef
type jsiiProxy_IDeploymentStrategyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDeploymentStrategyRef) DeploymentStrategyRef() *DeploymentStrategyReference {
	var returns *DeploymentStrategyReference
	_jsii_.Get(
		j,
		"deploymentStrategyRef",
		&returns,
	)
	return returns
}


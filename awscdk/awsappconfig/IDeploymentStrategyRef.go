package awsappconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeploymentStrategy.
// Experimental.
type IDeploymentStrategyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DeploymentStrategy resource.
	// Experimental.
	DeploymentStrategyRef() *DeploymentStrategyReference
}

// The jsii proxy for IDeploymentStrategyRef
type jsiiProxy_IDeploymentStrategyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IDeploymentStrategyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentStrategyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


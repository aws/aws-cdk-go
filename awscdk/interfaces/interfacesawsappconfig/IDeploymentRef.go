package interfacesawsappconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Deployment.
// Experimental.
type IDeploymentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Deployment resource.
	// Experimental.
	DeploymentRef() *DeploymentReference
}

// The jsii proxy for IDeploymentRef
type jsiiProxy_IDeploymentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDeploymentRef) DeploymentRef() *DeploymentReference {
	var returns *DeploymentReference
	_jsii_.Get(
		j,
		"deploymentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


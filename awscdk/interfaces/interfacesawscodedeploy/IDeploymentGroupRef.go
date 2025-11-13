package interfacesawscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodedeploy/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeploymentGroup.
// Experimental.
type IDeploymentGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DeploymentGroup resource.
	// Experimental.
	DeploymentGroupRef() *DeploymentGroupReference
}

// The jsii proxy for IDeploymentGroupRef
type jsiiProxy_IDeploymentGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDeploymentGroupRef) DeploymentGroupRef() *DeploymentGroupReference {
	var returns *DeploymentGroupReference
	_jsii_.Get(
		j,
		"deploymentGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


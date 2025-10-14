package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeploymentGroup.
// Experimental.
type IDeploymentGroupRef interface {
	constructs.IConstruct
	// A reference to a DeploymentGroup resource.
	// Experimental.
	DeploymentGroupRef() *DeploymentGroupReference
}

// The jsii proxy for IDeploymentGroupRef
type jsiiProxy_IDeploymentGroupRef struct {
	internal.Type__constructsIConstruct
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


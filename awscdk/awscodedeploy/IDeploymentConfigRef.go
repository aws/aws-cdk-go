package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeploymentConfig.
// Experimental.
type IDeploymentConfigRef interface {
	constructs.IConstruct
	// A reference to a DeploymentConfig resource.
	// Experimental.
	DeploymentConfigRef() *DeploymentConfigReference
}

// The jsii proxy for IDeploymentConfigRef
type jsiiProxy_IDeploymentConfigRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDeploymentConfigRef) DeploymentConfigRef() *DeploymentConfigReference {
	var returns *DeploymentConfigReference
	_jsii_.Get(
		j,
		"deploymentConfigRef",
		&returns,
	)
	return returns
}


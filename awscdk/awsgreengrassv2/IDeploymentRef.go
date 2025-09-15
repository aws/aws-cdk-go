package awsgreengrassv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrassv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Deployment.
// Experimental.
type IDeploymentRef interface {
	constructs.IConstruct
	// A reference to a Deployment resource.
	// Experimental.
	DeploymentRef() *DeploymentReference
}

// The jsii proxy for IDeploymentRef
type jsiiProxy_IDeploymentRef struct {
	internal.Type__constructsIConstruct
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


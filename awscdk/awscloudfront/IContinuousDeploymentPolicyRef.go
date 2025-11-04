package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContinuousDeploymentPolicy.
// Experimental.
type IContinuousDeploymentPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ContinuousDeploymentPolicy resource.
	// Experimental.
	ContinuousDeploymentPolicyRef() *ContinuousDeploymentPolicyReference
}

// The jsii proxy for IContinuousDeploymentPolicyRef
type jsiiProxy_IContinuousDeploymentPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IContinuousDeploymentPolicyRef) ContinuousDeploymentPolicyRef() *ContinuousDeploymentPolicyReference {
	var returns *ContinuousDeploymentPolicyReference
	_jsii_.Get(
		j,
		"continuousDeploymentPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContinuousDeploymentPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContinuousDeploymentPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


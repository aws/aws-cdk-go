package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContinuousDeploymentPolicy.
// Experimental.
type IContinuousDeploymentPolicyRef interface {
	constructs.IConstruct
	// A reference to a ContinuousDeploymentPolicy resource.
	// Experimental.
	ContinuousDeploymentPolicyRef() *ContinuousDeploymentPolicyReference
}

// The jsii proxy for IContinuousDeploymentPolicyRef
type jsiiProxy_IContinuousDeploymentPolicyRef struct {
	internal.Type__constructsIConstruct
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


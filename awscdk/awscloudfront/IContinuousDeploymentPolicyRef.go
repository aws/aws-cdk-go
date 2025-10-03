package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContinuousDeploymentPolicy.
// Experimental.
type IContinuousDeploymentPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IContinuousDeploymentPolicyRef
type jsiiProxy_IContinuousDeploymentPolicyRef struct {
	internal.Type__constructsIConstruct
}


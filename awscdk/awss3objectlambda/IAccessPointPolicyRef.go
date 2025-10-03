package awss3objectlambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3objectlambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessPointPolicy.
// Experimental.
type IAccessPointPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAccessPointPolicyRef
type jsiiProxy_IAccessPointPolicyRef struct {
	internal.Type__constructsIConstruct
}


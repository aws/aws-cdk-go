package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceProfile.
// Experimental.
type IInstanceProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInstanceProfileRef
type jsiiProxy_IInstanceProfileRef struct {
	internal.Type__constructsIConstruct
}


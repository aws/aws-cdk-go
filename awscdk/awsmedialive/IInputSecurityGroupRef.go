package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InputSecurityGroup.
// Experimental.
type IInputSecurityGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInputSecurityGroupRef
type jsiiProxy_IInputSecurityGroupRef struct {
	internal.Type__constructsIConstruct
}


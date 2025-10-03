package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityKey.
// Experimental.
type ISecurityKeyRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISecurityKeyRef
type jsiiProxy_ISecurityKeyRef struct {
	internal.Type__constructsIConstruct
}


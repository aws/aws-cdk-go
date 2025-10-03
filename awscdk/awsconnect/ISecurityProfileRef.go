package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityProfile.
// Experimental.
type ISecurityProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISecurityProfileRef
type jsiiProxy_ISecurityProfileRef struct {
	internal.Type__constructsIConstruct
}


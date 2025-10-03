package awsopensearchserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityConfig.
// Experimental.
type ISecurityConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISecurityConfigRef
type jsiiProxy_ISecurityConfigRef struct {
	internal.Type__constructsIConstruct
}


package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Secret.
// Experimental.
type ISecretRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISecretRef
type jsiiProxy_ISecretRef struct {
	internal.Type__constructsIConstruct
}


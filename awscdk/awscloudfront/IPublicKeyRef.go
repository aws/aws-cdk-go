package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PublicKey.
// Experimental.
type IPublicKeyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPublicKeyRef
type jsiiProxy_IPublicKeyRef struct {
	internal.Type__constructsIConstruct
}


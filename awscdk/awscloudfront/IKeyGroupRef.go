package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a KeyGroup.
// Experimental.
type IKeyGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IKeyGroupRef
type jsiiProxy_IKeyGroupRef struct {
	internal.Type__constructsIConstruct
}


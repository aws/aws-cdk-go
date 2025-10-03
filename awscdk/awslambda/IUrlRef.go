package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Url.
// Experimental.
type IUrlRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUrlRef
type jsiiProxy_IUrlRef struct {
	internal.Type__constructsIConstruct
}

